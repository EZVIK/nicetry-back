package main

import "fmt"

// Weekday 为枚举定义一个类型
type Weekday uint32

// 如果不关心枚举值的具体数值， 可以用iota来简化定义
const (
	Unknown Weekday = iota // 为第一个枚举之定义为Unknown 值为 0 表示未确定 或者 未选中
	Monday
	Tuesday
	Wednesday
	Thursday
	Firday
	Saturday
	Sunday
	end // 验证合法性 不使用
)

func (w Weekday) IsValid() bool {
	return w < end
}

func main() {
	fmt.Println(Monday.IsValid())
	fmt.Println(Weekday(2).IsValid())
	fmt.Println(Weekday(100).IsValid())
}
