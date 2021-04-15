package app



type Response struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
}


func NewRes(data interface{}) Response {
	return Response{Code: 200, Msg: "Success", Data: data}
}

func NewErrRes(code int, msg string, err string) Response {
	return Response{Code: code, Msg: msg, Error: err}
}

