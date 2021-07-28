package main

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

const (
	SecretID  = "AKIDmMfikz6pVw83i8JJYS7dqVuYx53lOByV"
	SecretKey = "sQRRL7XeU3JMAh5ExHLipfFfCcssma4H"
)

func main() {

	//println(strconv.FormatUint(0, 1))

	// 将 examplebucket-1250000000 和 COS_REGION修改为真实的信息
	u, _ := url.Parse("https://nne259o445k7v88a-1257958891.cos.ap-shanghai.myqcloud.com")

	// 用于Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, _ := url.Parse("https://service.cos.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}

	// 1.永久密钥
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SecretID,
			SecretKey: SecretKey,
		},
	})

	// 查询桶列表
	//s, _, err := c.Service.Get(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//for _, b := range s.Buckets {
	//	fmt.Printf("%#v\n", b)
	//}

	//name := "test/objectPut.go"
	//// 1.通过字符串上传对象
	//f := strings.NewReader("test")
	//_, err := c.Object.Put(context.Background(), name, f, nil)
	//if err != nil {
	//	panic(err)
	//}

	// 2.通过本地文件上传对象
	// name: 上传目录名称  filePath: local file path
	//_, err := c.Object.PutFromFile(context.Background(), "test/t.go", "./test/t.go", nil)
	//if err != nil {
	//	panic(err)
	//}

	// 下载文件
	name := "test/t.go"
	_, err := c.Object.GetToFile(context.Background(), name, "test.go", nil)
	if err != nil {
		panic(err)
	}

}
