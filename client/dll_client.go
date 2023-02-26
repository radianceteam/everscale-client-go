//nolint:godot
package client

import (
	"encoding/json"
	"errors"
	"unsafe"

	"go.uber.org/zap"
)

// #cgo CFLAGS: -I${SRCDIR}
// #include <stdlib.h>
// #include "tonclient.h"
//
// void callbackProxy(
// 		uint32_t request_id,
//		tc_string_data_t params_json,
//		uint32_t response_type,
//		bool finished);
//
// static void call_tc_request(
//		uint32_t context,
//		tc_string_data_t name,
//		tc_string_data_t params_json,
//		uint32_t request_id
// ) {
//		tc_request(context, name, params_json, request_id, callbackProxy);
//		free((void*)name.content);
//		free((void*)params_json.content);
// }
//
import "C"

var ErrContextIsClosed = errors.New("context is closed")

type ResponseCode uint32

const (
	ResponseCodeSuccess    = ResponseCode(C.tc_response_success)
	ResponseCodeError      = ResponseCode(C.tc_response_error)
	ResponseCodeNop        = ResponseCode(C.tc_response_nop) // User-level code never sees this event
	ResponseCodeAppRequest = ResponseCode(C.tc_response_app_request)
	ResponseCodeAppNotify  = ResponseCode(C.tc_response_app_notify)
	ResponseCodeCustom     = ResponseCode(C.tc_response_custom)
)

func newTcStr(str []byte) C.tc_string_data_t {
	return C.tc_string_data_t{
		len:     C.uint32_t(len(str)),
		content: C.CString(string(str)),
	}
}

func newBytesFromTcStr(data C.tc_string_data_t) []byte {
	if data.len == 0 {
		return nil
	}

	return C.GoBytes(unsafe.Pointer(data.content), C.int(data.len)) // nolint nlreturn
}

func newDLLResponse(rawBytes []byte, responseType ResponseCode) *RawResponse {
	res := &RawResponse{
		Code: responseType,
	}
	if responseType == ResponseCodeError {
		var sdkErr Error
		err := json.Unmarshal(rawBytes, &sdkErr)
		if err != nil {
			res.Error = err
		} else {
			res.Error = &sdkErr
		}
	} else {
		res.Data = rawBytes
	}

	return res
}

//export callbackProxy
func callbackProxy(requestIDRaw C.uint32_t, data C.tc_string_data_t, responseTypeRaw C.uint32_t, finishedRaw C.bool) {
	requestID := uint32(requestIDRaw)
	finished := bool(finishedRaw)
	responseType := ResponseCode(responseTypeRaw)
	rawBody := newBytesFromTcStr(data)
	zap.L().Debug("new response",
		zap.Uint32("request_id", requestID),
		zap.Uint32("response_type", uint32(responseType)),
		zap.Bool("finished", finished),
		zap.ByteString("body", rawBody))
	responses, closeSignal, isFound := globalMultiplexer.GetChannels(requestID, finished)
	if !isFound {
		// ignore not found maybe some will be send after close
		return
	}

	if responseType == ResponseCodeNop {
		if finished {
			close(responses)
		}

		return
	}

	select {
	case responses <- newDLLResponse(rawBody, responseType):
		if finished {
			close(responses)
		}
	case <-closeSignal:
		close(responses)
		globalMultiplexer.DeleteByRequestID(requestID)
	}
}

type RawResponse struct {
	Data  []byte
	Code  ResponseCode
	Error error
}

type dllClient interface {
	waitErrorOrResult(method string, body interface{}) ([]byte, error)
	waitErrorOrResultUnmarshal(method string, body interface{}, dst interface{}) error
	resultsChannel(method string, body interface{}) (<-chan *RawResponse, error)
	close()
}

func newDLLClient(rawConfig []byte, maxCGOConcurrentThreads uint) (*dllClientCtx, error) {
	c := &dllClientCtx{
		closeSignal:             make(chan struct{}),
		maxCGOConcurrentThreads: make(chan struct{}, maxCGOConcurrentThreads),
	}

	for i := uint(0); i < maxCGOConcurrentThreads; i++ {
		c.maxCGOConcurrentThreads <- struct{}{}
	}

	return c, c.createContext(rawConfig)
}

type dllClientCtx struct {
	closeSignal             chan struct{}
	maxCGOConcurrentThreads chan struct{}
	ctx                     C.uint32_t
}

type contextCreateResponse struct {
	Result uint32 `json:"result"`
	Error  *Error `json:"error"`
}

func (c *dllClientCtx) createContext(data []byte) error {
	cStringData := newTcStr(data)
	rawHandler := C.tc_create_context(cStringData)
	C.free(unsafe.Pointer(cStringData.content))

	defer C.tc_destroy_string(rawHandler)
	rawResponse := newBytesFromTcStr(C.tc_read_string(rawHandler))
	var response contextCreateResponse
	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return response.Error
	}
	c.ctx = C.uint32_t(response.Result)

	return nil
}

func (c *dllClientCtx) close() {
	C.tc_destroy_context(c.ctx)
	close(c.closeSignal)
}

func (c *dllClientCtx) waitErrorOrResultUnmarshal(method string, body interface{}, dst interface{}) error {
	rawData, err := c.waitErrorOrResult(method, body)
	if err != nil {
		return err
	}

	return json.Unmarshal(rawData, dst)
}

func (c *dllClientCtx) resultsChannel(method string, body interface{}) (<-chan *RawResponse, error) {
	var rawBody []byte
	if body != nil {
		var err error
		rawBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	responses := make(chan *RawResponse, 1) // need 1 because of deadlock when async implemented in sync way
	requestID := globalMultiplexer.SetChannels(responses, c.closeSignal)
	zap.L().Debug("new request",
		zap.Uint32("request_id", requestID),
		zap.String("method", method),
		zap.ByteString("body", rawBody))
	turn := <-c.maxCGOConcurrentThreads
	defer func() {
		c.maxCGOConcurrentThreads <- turn
	}()
	C.call_tc_request(c.ctx, newTcStr([]byte(method)), newTcStr(rawBody), C.uint32_t(requestID))

	return responses, nil
}

func (c *dllClientCtx) waitErrorOrResult(method string, body interface{}) ([]byte, error) {
	responses, err := c.resultsChannel(method, body)
	if err != nil {
		return nil, err
	}
	var data []byte

	for {
		select {
		case r, ok := <-responses:
			if !ok {
				return data, err
			}
			if r.Error != nil && err == nil {
				err = r.Error
			}
			if r.Data != nil && data == nil {
				data = r.Data
			}
		case <-c.closeSignal:
			return nil, ErrContextIsClosed
		}
	}
}
