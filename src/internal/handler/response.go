package handler

import "gin-web-template/internal/errors"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func NewResponse(code int, msg string, data interface{}) Response {
	return Response{Code: code, Msg: msg, Data: data}
}

func NewSuccessResponse(data interface{}) Response {
	return NewResponse(0, "Success", data)
}

func NewErrorResponse(err errors.SrvErr) Response {
	return NewResponse(err.Code(), err.Message(), nil)
}
