package client

import (
	"encoding/json"
	"fmt"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -I${SRCDIR}
// #include "tonclient.h"
//
// void cb_proxy(
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
// 		tc_request(context, name, params_json, request_id, cb_proxy);
// }
//
import "C"

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

//export cb_proxy
func cb_proxy(id C.uint32_t, data C.tc_string_data_t, responseType C.uint32_t, finished C.bool) {
	lock.Lock()
	result, isFound := callbacks[id]
	if bool(finished) && isFound {
		fmt.Println("cb_proxy", finished)
		delete(callbacks, id)
		defer close(result)
	}
	lock.Unlock()
	if !isFound {
		fmt.Println("cb_proxy no channel to send", id)

		return
	}
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
	result <- res
}

type dllResponse struct {
	Data  []byte
	Code  int
	Error error
}

var (
	lock             sync.Locker = &sync.Mutex{}
	requestIDCounter uint32
	callbacks        = make(map[C.uint32_t]chan *dllResponse)
)

type dllClient interface {
	Request(method string, body []byte) <-chan *dllResponse
	Close()
}

func newDLLClient(rawConfig []byte) (dllClient, error) {
	c := &dllClientCtx{}

	return c, c.createContext(rawConfig)
}

type dllClientCtx struct {
	ctx C.uint32_t
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

func (c *dllClientCtx) Request(method string, body []byte) <-chan *dllResponse {
	responses := make(chan *dllResponse, 1) // need 1 because of deadlock when async implemented in sync way

	lock.Lock()
	requestIDCounter++
	id := C.uint32_t(requestIDCounter)
	callbacks[id] = responses
	lock.Unlock()

	C.call_tc_request(c.ctx, newTcStr([]byte(method)), newTcStr(body), id)

	return responses
}

func (c *dllClientCtx) Close() {
	C.tc_destroy_context(c.ctx)
}

func getFirstErrorOrResult(responses <-chan *dllResponse) ([]byte, error) {
	var err error
	var data []byte
	for response := range responses {
		if response.Error != nil && err == nil {
			err = response.Error
		}
		if response.Data != nil && data == nil {
			data = response.Data
		}
	}

	return data, err
}
