package client

import (
	"fmt"
)

var errorCodesToErrorTypes = map[uint32]string{}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	return fmt.Sprintf("sdk_error_code=%d,\n sdk_error_code_description=%s,\n sdk_error_msg='%s',\n sdk_error_data: '%s'",
		e.Code, errorCodesToErrorTypes[e.Code], e.Message, e.Data)
}
