package main

import (
	"fmt"
	"github.com/guonaihong/gout"
	"time"
)

// 用于解析 服务端 返回的http body
type RspBody struct {
	ErrMsg  string        `json:"msg"`
	ErrCode int           `json:"code"`
	Data    LoginResponse `json:"data"`
}

type LoginResponse struct {
	JWT      string   `json:"jwt"`
	UserInfo UserInfo `json:"user_info"`
}

type UserInfo struct {
	ID            int    `json:"ID"`
	avatar        string `json:"avatar"`
	desc          string `json:"desc"`
	link          int    `json:"link"`
	mail          string `json:"mail"`
	nickname      string `json:"nickname"`
	password      string `json:"password"`
	points        int    `json:"points"`
	recommend_by  int    `json:"recommend_by"`
	enable_status bool   `json:"enable_status"`
	CreatedAt     string `json:"CreatedAt"`
	UpdatedAt     string `json:"UpdatedAt"`
	DeletedAt     string `json:"DeletedAt"`
}

// 用于解析 服务端 返回的http header
type RspHeader struct {
	Sid  string `header:"sid"`
	Time int    `header:"time"`
}

func (r *RspBody) String() string {
	return r.Data.JWT
}

func main() {
	rsp := RspBody{}
	header := RspHeader{}

	//code := 0
	err := gout.

		// POST请求
		POST("http://127.0.0.1:6709/api/v1/user/login").

		// 打开debug模式
		Debug(true).

		// 设置查询字符串
		//SetQuery(gout.H{"page": 10, "size": 10}).

		// 设置http header
		SetHeader(gout.H{"X-IP": "127.0.0.1", "sid": fmt.Sprintf("%x", time.Now().UnixNano())}).

		// SetJSON设置http body为json
		// 同类函数有SetBody, SetYAML, SetXML, SetForm, SetWWWForm
		SetJSON(gout.H{
			"mail":     "eee@ee.com",
			"password": "sonofbitch",
		}).

		// BindJSON解析返回的body内容
		// 同类函数有BindBody, BindYAML, BindXML
		BindJSON(&rsp).

		// 解析返回的http header
		BindHeader(&header).
		// http code
		// Code(&code).

		// 结束函数
		Do()

	// 判断错误
	if err != nil {
		fmt.Printf("send fail:%s\n", err)
	}

	fmt.Println(rsp.String())
}

/*
> POST /?page=10&size=10 HTTP/1.1
> Sid: 15d9b742ef32c130
> X-Ip: 127.0.0.1
> Content-Type: application/json
>

{
   "text": "gout"
}


*/
