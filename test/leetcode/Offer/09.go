package main

import (
	"container/list"
	"fmt"
)

/**
用两个栈实现一个队列。

队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，

分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

*/

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

// stack first in last  out

// queue first in first out

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor(init []int) CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {

	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}

	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}

	return -1
}

func main() {

	stack1 := *list.New()

	stack1.PushBack(999)
	stack1.PushFront(1)
	stack1.PushBack(998)

	fmt.Println(stack1.Front().Value)
}
