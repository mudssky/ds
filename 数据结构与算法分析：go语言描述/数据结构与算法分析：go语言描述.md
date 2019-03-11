# 数据结构与算法分析：Go语言描述

手头正好有一本书，《数据结构与算法分析：c++语言描述》，当年学数据结构的时候，真的完全没花心思进去，学出来跟没学一样，所以现在准备重学。

那些数据结构，全部用go语言重新实现一遍。正好go语言里面其实也没有这方面的标准库，写完了，作为一个自己的库使用。



前五章，200多页内容都是讲c++的语法，面向对象之类的。第6章开始才是抽象数据结构。所以前面的内容暂且不看。

因为原书是c++语言的，所以我干脆先用c++实现一遍？毕竟大部分很大部分内容都是c++相关。

可以先c++看书实现一遍，然后go语言不看书实现一遍。

试了一下还是放弃了，我再次感受到以前学习c++时的痛苦，从头到尾把前面5章概览了一边，c++繁多的特性，再次出现在我的脑海中。首先是语法繁琐，代码分为.h头文件和cpp文件，函数声明和定义加起来要写两遍。

其次是ide不好用，vs智能提示很弱，配置麻烦（我不会配置啊，正常不应该开箱即用吗），还有各种莫名其妙的编译器错误

搞了半天把第一个基于静态数组的列表实现并测试完，我就想放弃了。垃圾回收的部分还没写过，不过应该也是比较繁琐的。

所以说直接用golang实现吧。

## 1.列表

### 01.用切片slice设计一个ArrayList

go 语言里面，其实这么搞没有太大必要，因为slice本来功能已经足够强大了。这边优势是，元素类型用interface{}，所以各种类型都可以放进去。

```go
// Package arraylist 基于go的基础类型切片封装一个arraylist
package arraylist

import "fmt"

// ListType 用来确定元素类型,interface{}类型可以接受各种类型
type ListType interface{}

const (
	// INITCAP 初始容量
	INITCAP = 10
)

// ArrayList 基于go切片类型设计的列表
type ArrayList struct {
	list []ListType
	len  int //元素个数
	cap  int //底层切片容量大小
}

// Init ArrayList的初始化方法，列表初始化大小为10
func (a *ArrayList) Init() *ArrayList {
	a.list = make([]ListType, INITCAP)
	a.cap = INITCAP
	a.len = 0
	return a
}

// New 相当于一个工厂方法，调用这个方法，直接返回一个初始化好的ArrayList指针
func New() *ArrayList {
	list := new(ArrayList)
	return list.Init()
}

// GetLen 获取ArrayList当前的元素个数
func (a *ArrayList) GetLen() int {
	return a.len

}

func (a *ArrayList) expandCap() {
	a.cap *= 2
	// 创建一个新的切片，每次容量变为两倍
	newist := make([]ListType, a.cap)
	// 将旧list数据全部拷贝到新list
	copy(newist[:a.len], a.list[:a.len])

	a.list = newist
}

// Append 在ArrayList列表尾部添加一个元素
// 如果列表容量不够，则会扩容到两倍，并把数据拷贝过去
func (a *ArrayList) Append(item ListType) {
	if a.len == a.cap {
		a.expandCap()
		a.list[a.len] = item
		a.len++
	} else {
		a.list[a.len] = item
		a.len++
	}
}

// Push 在ArrayList列表尾部添加一个元素，使用Append的实现
func (a *ArrayList) Push(item ListType) {
	a.Append(item)
}

// Pop 返回列表最后一个元素，列表长度减一
func (a *ArrayList) Pop() ListType {
	if a.len == 0 {
		panic("pop error , no element in the list")
	}
	a.len--
	return a.list[a.len]
}

// Empty 判断ArrayList里的列表是否为空
func (a *ArrayList) Empty() bool {
	if a.len == 0 {
		return true
	}
	return false
}

// Insert 在指定位置pos，插入一个数据 item
// 也可以支持在后面插入，相当于实现Append的功能
// 如果给出一个超出范围的数字，会触发panic
func (a *ArrayList) Insert(item ListType, pos int) {
	if a.len == a.cap {
		a.expandCap()
	}
	if pos < 0 || pos > a.len {
		panic("insert error, pos is out of range")
	}
	for i := a.len; i > pos; i-- {
		a.list[i] = a.list[i-1]
	}
	a.list[pos] = item
	a.len++
}

// Erase 擦除指定位置pos的数据
func (a *ArrayList) Erase(pos int) {
	if pos < 0 || pos > a.len {
		panic("erase error , pos is out of range ")
	}
	for i := pos; i < a.len; i++ {
		a.list[i] = a.list[i+1]
	}
	a.len--
}

// Show 将列表打印输出
func (a *ArrayList) Show() {
	for i := 0; i < a.len; i++ {
		fmt.Print(a.list[i], " ")
	}
	fmt.Println()
}

```



## 02.链表

在arraylist中元素被排列成一个线性序列，当添加的数据量达到容量值的时候，要进行拷贝。

插入的时候也会进行频繁的位移和拷贝，导致效率低下。

链表是由成为节点的元素组成的序列，其中每个元素包括两部分，

1. data存放列表的一个元素
2. next，存放一个指针，指出下一个列表元素的节点的位置，如果没有下一个元素，将用一个特殊的nil值

另外，我们需要记录第一个节点的位置，这样我们可以通过这个节点不断向后遍历到链表的所有节点。

