package stack

import (
	"fmt"
)

const (
	// INITCAP 初始容量
	INITCAP = 10
)

// Stack 栈类的结构体
// cap为底层list的容量，top表示栈顶下标，同时也是栈的元素个数
type Stack struct {
	list []interface{}
	cap  int
	top  int
}

// Init 初始化栈类，或者清空栈类
func (s *Stack) Init() *Stack {
	s.list = make([]interface{}, INITCAP)
	s.cap = INITCAP
	s.top = 0
	return s
}

// New 返回一个初始化好的栈
func New() *Stack {
	return new(Stack).Init()
}

// Empty 通过判断栈顶下标的位置是否为0，判断栈是否为空
func (s *Stack) Empty() bool {
	if s.top == 0 {
		return true
	}
	return false
}

// 当栈类里面的数组容量不够时，调用这个方法扩容到两倍，原来的数据也会拷贝过去
func (s *Stack) expandCap() {
	s.cap *= 2
	// 创建一个新的切片，每次容量变为两倍
	newist := make([]interface{}, s.cap)
	// 将旧list数据全部拷贝到新list
	copy(newist[:s.top], s.list[:s.top])
	s.list = newist
}

// Push 添加元素到栈尾
func (s *Stack) Push(val interface{}) {
	if s.top == s.cap {
		s.expandCap()
	}
	s.list[s.top] = val
	s.top++
}

// Pop 删除栈顶的一个元素并返回，没有添加错误处理，执行这个操作是要先判断栈是否为空
func (s *Stack) Pop() interface{} {
	s.top--
	return s.list[s.top]
}

// Top 返回栈顶的元素，如果栈为空，返回nil
func (s *Stack) Top() interface{} {
	if s.top == 0 {
		return nil
	}
	return s.list[s.top-1]
}

// Display 输出栈中的所有元素，栈底到栈顶，从左到右
func (s *Stack) Display() {
	for i := 0; i < s.top; i++ {
		fmt.Print(s.list[i], " ")
	}
	fmt.Println()
}
