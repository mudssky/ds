package linklist

import "fmt"

// Node 单链表的节点
type Node struct {
	Element interface{}
	Next    *Node
}

// LinkList 单链表类的结构体
type LinkList struct {
	Head *Node
	len  int
}

// Init 初始化linklist或者清空linklist
func (l *LinkList) Init() *LinkList {
	l.Head = nil
	l.len = 0
	return l
}

// New 返回一个初始化的linklist
func New() *LinkList {
	return new(LinkList).Init()
}

// Empty 通过头节点是否为空判断linklist是否为空
func (l *LinkList) Empty() bool {
	if l.Head == nil {
		return true
	}
	return false
}

// Len 获取linklist的节点数目
func (l *LinkList) Len() int { return l.len }

// Insert 插入某个值，在at节点的后面
func (l *LinkList) Insert(val interface{}, at *Node) {
	at.Next.Element = val
	l.len++
}

// Append 添加元素在链表的尾部
func (l *LinkList) Append(val interface{}) {
	if l.Empty() {
		l.Head = &Node{val, nil}
	} else {
		tmp := l.Head
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = &Node{val, nil}
	}
	l.len++
}

// Remove 移除linklist的某个节点
func (l *LinkList) Remove(at *Node) {
	if at == nil {
		return
	}
	if at == l.Head {
		l.Head = l.Head.Next
	} else {
		tmp := l.Head
		for tmp.Next != nil {
			if tmp.Next == at {
				tmp.Next = at.Next
			}
		}
	}
	l.len--
}

// Traversal 遍历单链表
func (l *LinkList) Traversal() {
	tmp := l.Head
	fmt.Print(tmp.Element)
	for tmp.Next != nil {
		tmp = tmp.Next
		fmt.Print(",", tmp.Element)
	}
	fmt.Println()
}
