package main

import (
	"fmt"
	"nicetry/pkg/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {

	for i := 0; i < 100; i++ {
		GetFakeNumber("你真123好")
	}
}

func GetFakeNumber(title string) {

	i := strconv.Itoa(int(time.Now().UnixNano()))

	res := utils.Md5Crypt(i, title)

	freg := regexp.MustCompile(`[0-9]`)

	params := freg.FindAllString(res, -1)

	start, end := 0, 6
	if params[0] == "0" {
		start = 1
		end += 1
	}
	fmt.Println(res)
	fmt.Println(strings.Join(params[start:end], ""))

}
