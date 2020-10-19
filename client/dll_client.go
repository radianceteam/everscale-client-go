package client

import (
	"encoding/json"
	"errors"
	"unsafe"
)

// #cgo CFLAGS: -I${SRCDIR}
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
// 		tc_request(context, name, params_json, request_id, callbackProxy);
// }
//
import "C"

var ErrContextIsClosed = errors.New("context is closed")

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

func newDLLResponse(data C.tc_string_data_t, responseType C.uint32_t) *dllResponse {
	rawBytes := newBytesFromTcStr(data)
	res := &dllResponse{
		Code: int(responseType),
	}
	if responseType == C.tc_response_error {
		var sdkErr ErrorSDK
		err := json.Unmarshal(rawBytes, &sdkErr)
		if err != nil {
			res.Error = err
		} else {
			res.Error = &sdkErr
		}
	} else if responseType == C.tc_response_success {
		res.Data = rawBytes
	}

	return res
}

//export callbackProxy
func callbackProxy(requestID C.uint32_t, data C.tc_string_data_t, responseType C.uint32_t, finished C.bool) {
	responses, closeSignal, isFound := globalMultiplexer.GetChannels(uint32(requestID), bool(finished))
	if !isFound {
		// ignore not found maybe some will be send after close
		return
	}

	select {
	case responses <- newDLLResponse(data, responseType):
		if bool(finished) {
			close(responses)
		}
	case <-closeSignal:
		close(responses)
		globalMultiplexer.DeleteByRequestID(uint32(requestID))
	}
}

type dllResponse struct {
	Data  []byte
	Code  int
	Error error
}

type dllClient interface {
	waitErrorOrResult(method string, body interface{}) ([]byte, error)
	waitErrorOrResultUnmarshal(method string, body interface{}, dst interface{}) error
	close()
}

func newDLLClient(rawConfig []byte) (dllClient, error) {
	c := &dllClientCtx{
		closeSignal: make(chan struct{}),
	}

	return c, c.createContext(rawConfig)
}

type dllClientCtx struct {
	closeSignal chan struct{}
	ctx         C.uint32_t
}

type contextCreateResponse struct {
	Result uint32    `json:"result"`
	Error  *ErrorSDK `json:"error"`
}

func (c *dllClientCtx) createContext(data []byte) error {
	rawHandler := C.tc_create_context(newTcStr(data))
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
	// fmt.Println(string(rawData))
	return json.Unmarshal(rawData, dst)
}

func (c *dllClientCtx) waitErrorOrResult(method string, body interface{}) ([]byte, error) {
	var rawBody []byte
	if body != nil {
		var err error
		rawBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	responses := make(chan *dllResponse, 1) // need 1 because of deadlock when async implemented in sync way
	requestID := globalMultiplexer.SetChannels(responses, c.closeSignal)
	// TODO maybe add global worker pool later for CGO calls
	C.call_tc_request(c.ctx, newTcStr([]byte(method)), newTcStr(rawBody), C.uint32_t(requestID))
	var (
		err  error
		data []byte
	)

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
