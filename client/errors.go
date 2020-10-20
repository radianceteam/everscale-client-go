package client

import (
	"fmt"
)

type ErrorSDK struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *ErrorSDK) Error() string {
	if e == nil {
		return ""
	}

	return fmt.Sprintf("%s, sdk_error_code=%d", e.Message, e.Code)
}
