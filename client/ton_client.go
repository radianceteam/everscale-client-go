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
//void cb_proxy(
//uint32_t request_id,
//tc_string_data_t params_json,
//uint32_t response_type,
//bool finished);
//
// static void call_tc_request(
//		uint32_t context,
//		tc_string_data_t name,
//		tc_string_data_t params_json,
//		uint32_t request_id
// ) {
// 		tc_request(context, name, params_json, request_id, cb_proxy);
// }
import "C"

func NewTcStr(str []byte) C.tc_string_data_t {
	return C.tc_string_data_t{
		len:     C.uint32_t(len(str)),
		content: C.CString(string(str)),
	}
}

func NewBytesFromTcStr(data C.tc_string_data_t) []byte {
	if data.len == 0 {
		return nil
	}

	return C.GoBytes(unsafe.Pointer(data.content), C.int(data.len))
}

//export cb_proxy
func cb_proxy(id C.uint32_t, data C.tc_string_data_t, responseType C.uint32_t, finished C.bool) {
	lock.Lock()
	result := callbacks[id]
	delete(callbacks, id)
	lock.Unlock()
	rawBytes := NewBytesFromTcStr(data)
	res := &response{
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

type response struct {
	Data  []byte
	Code  int
	Error error
}

var lock sync.Mutex
var requestIdCounter uint32
var callbacks map[C.uint32_t]chan *response

func init() {
	callbacks = make(map[C.uint32_t]chan *response)
}

type dllClientCtx struct {
	ctx     C.uint32_t
	counter C.uint
}

type contextCreateResponse struct {
	Result uint32    `json:"result"`
	Error  *ErrorSDK `json:"error"`
}

func (c *dllClientCtx) createContext(data []byte) error {
	rawResponse := NewBytesFromTcStr(C.tc_read_string(C.tc_create_context(NewTcStr(data))))
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

func (c *dllClientCtx) Request(method string, body []byte) ([]byte, error) {
	result := make(chan *response, 1) // need 1 because of deadlock when async implemented in sync way

	lock.Lock()
	requestIdCounter++
	id := C.uint32_t(requestIdCounter)
	callbacks[id] = result
	lock.Unlock()

	C.call_tc_request(c.ctx, NewTcStr([]byte(method)), NewTcStr(body), id)

	response := <-result
	if response.Error != nil {
		return nil, response.Error
	}
	fmt.Println(method, string(response.Data))
	return response.Data, nil
}

func (c *dllClientCtx) Close() {
	C.tc_destroy_context(c.ctx)
}

type dllClient interface {
	Request(method string, body []byte) ([]byte, error)
	Close()
}

func NewDLLClient(rawConfig []byte) (dllClient, error) {
	c := &dllClientCtx{}
	err := c.createContext(rawConfig)
	return c, err
}
