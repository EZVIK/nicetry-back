package utils

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

//func SplicingStringArray(element interface{}, key string) (string, error) {
//
//	etype := element.([]interface{})
//
//	var buffer bytes.Buffer
//	// 合并标签 为字符串
//	for FileOp, k := range etype {
//		buffer.WriteString(k.key.(string))
//		if FileOp != len(etype)-1 {
//			buffer.WriteString(" ")
//		}
//	}
//	return buffer.String(), nil
//}

func GetNoNumber(title string) string {

	i := strconv.Itoa(int(time.Now().UnixNano()))

	res := Md5Crypt(i, title)

	freg := regexp.MustCompile(`[0-9]`)

	params := freg.FindAllString(res, -1)

	start, end := 0, 6

	if params[0] == "0" {
		start = 1
		end += 1
	}

	return strings.Join(params[start:end], "")
}
