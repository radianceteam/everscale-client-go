package client

import "strconv"

type ErrorSDK struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func (e *ErrorSDK) Error() string {
	if e == nil {
		return ""
	}
	return strconv.Itoa(e.Code) + " " +e.Message
}
