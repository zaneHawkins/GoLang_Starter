package handler

import (
	T "src/types"
)

func ServiceError(code int, msg string, err error) *T.ServiceError {
	return &T.ServiceError{
		Code:    code,
		Message: msg,
		Error:   err,
	}
}
