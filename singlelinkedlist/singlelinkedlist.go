package singlelinkedlist

import "fmt"

// ListNode 单向链表的节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// List 单向链表的节点封装
type List struct {
	head *ListNode //记录头节点的位置
	len  int       //记录链表中的元素个数
}

// Init 单向链表的初始化方法
func (l *List) Init() *List {
	l.head.Val = 0
	l.head.Next = nil
	l.len = 0
	return l
}

// New 创建链表并进行初始化操作
func New() *List { return new(List).Init() }

// Len 返回链表的长度
func (l *List) Len() int { return l.len }

// Append  在链表尾部追加元素
func (l *List) Append(val int) {
	tmp := l.head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &ListNode{val, nil}
	l.len++
}

// Traversal 遍历链表并且打印出来
func (l *List) Traversal() {
	tmp := l.head
	fmt.Print(tmp.Val)
	for tmp.Next != nil {
		tmp = tmp.Next
		fmt.Print(",", tmp.Val)
	}
	fmt.Println()
}

// Add 在单向链表的尾部添加节点
func (ln *ListNode) Add(val int) {
	for ln.Next != nil {
		ln = ln.Next
	}
	ln.Next = &ListNode{val, nil}

}

// TraversalFrom 从单向链表的某个节点开始遍历
func (ln *ListNode) TraversalFrom(begin *ListNode) {
	fmt.Print(begin.Val)
	for begin.Next != nil {
		begin = begin.Next
		fmt.Print(",", begin.Val)
	}
	fmt.Println()

}

// Traversal 遍历单向链表，从给出的节点开始
func (ln *ListNode) Traversal() {
	tmp := ln
	fmt.Print(tmp.Val)
	for tmp.Next != nil {
		tmp = tmp.Next
		fmt.Print(",", tmp.Val)
	}
	fmt.Println()
}
