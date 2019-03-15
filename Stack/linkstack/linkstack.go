package linkstack

import "fmt"

// node 链式栈的节点
type node struct {
	value interface{}
	next  *node
}

// LinkStack 链式栈的结构体
type LinkStack struct {
	top *node
	len int
}

// Init 初始化链栈，或者情况链栈
func (l *LinkStack) Init() *LinkStack {
	l.top = nil
	l.len = 0
	return l
}

// New 返回一个初始化好的链栈
func New() *LinkStack {
	return new(LinkStack).Init()
}

// Empty 通过判断链栈顶部节点是否为nil，判断链栈是否为空栈
func (l *LinkStack) Empty() bool {
	if l.top == nil {
		return true
	}
	return false
}

// Top 返回链栈顶部节点的值
func (l *LinkStack) Top() interface{} {
	return l.top.value
}

// Push 在链栈的尾部添加元素
func (l *LinkStack) Push(val interface{}) {
	if l.top == nil {
		l.top = &node{val, nil}
	} else {
		newNode := &node{val, nil}
		newNode.next = l.top
		l.top = newNode
	}
	l.len++
}

// Pop 删除链栈顶部节点，并返回那个节点的值
func (l *LinkStack) Pop() interface{} {
	if l.top == nil {
		panic("error,no nodes in this linkstack")
	}
	topVal := l.top.value
	l.top = l.top.next
	l.len--
	return topVal

}

// Display 输出链栈的所有元素，从顶部到底部，从左到右
func (l *LinkStack) Display() {
	for l.top.next != nil {
		fmt.Print(l.top.value, " ")
		l.top = l.top.next
	}
	fmt.Println()
}
