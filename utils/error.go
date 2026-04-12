package utils

import (
	"net/http"

	"github.com/zikrykr/backend-test-desent/infrastructure"
)

type ErrorObj struct {
	HTTPStatus int    `json:"http_status"`
	Message    string `json:"message"`
	Detail     string `json:"detail"`
	Error      error  `json:"error"`
}

func NewError(logger *infrastructure.Logger, status int, message string, err error) *ErrorObj {
	logger.Error(message, err)
	return &ErrorObj{
		HTTPStatus: status,
		Message:    message,
		Detail:     http.StatusText(status),
		Error:      err,
	}
}

func NotFoundError(logger *infrastructure.Logger, message string, err error) *ErrorObj {
	return NewError(logger, http.StatusNotFound, message, err)
}

func BadRequestError(logger *infrastructure.Logger, message string, err error) *ErrorObj {
	return NewError(logger, http.StatusBadRequest, message, err)
}

func InternalServerError(logger *infrastructure.Logger, message string, err error) *ErrorObj {
	return NewError(logger, http.StatusInternalServerError, message, err)
}
