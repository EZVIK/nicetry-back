package app

import "nicetry/pkg/e"

type Response struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data,omitempty"`
}


func NewRes(data interface{}) Response {
	return Response{Code: 200, Msg: "Success", Data: data}
}


func NewErrRes(code int, msg string, data interface{}) Response {
	return Response{Code: code, Msg: msg, Data: data}
}


func NewErr(err *e.Error) Response {
	return Response{Code: err.StatusCode(), Msg: err.Msg(), Data: err.Details()}
}