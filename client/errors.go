package client

import (
	"fmt"
)

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	return fmt.Sprintf("%s, sdk_error_code=%d", e.Message, e.Code)
}
