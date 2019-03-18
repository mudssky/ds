[TOC]



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



### 02.简单链表

在arraylist中元素被排列成一个线性序列，当添加的数据量达到容量值的时候，要进行拷贝。

插入的时候也会进行频繁的位移和拷贝，导致效率低下。

链表是由成为节点的元素组成的序列，其中每个元素包括两部分，

1. data存放列表的一个元素
2. next，存放一个指针，指出下一个列表元素的节点的位置，如果没有下一个元素，将用一个特殊的nil值

另外，我们需要记录第一个节点的位置，这样我们可以通过这个节点不断向后遍历到链表的所有节点。



列表基本操作

1. 构造（Construction）

为了构造一个空列表，我们只需简单地把first设为空指针，就可以判断出列表是否为空。

2. 判空(Empty)

判断first是否为空，就可以判断出列表是否为空。

3. 遍历（Traverse）

遍历列表

4. 插入（Insertion）
5. 删除（Deletion）



基于go语言的实现

```go
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

```



以上是一个简单链表，除此之外，还有很多变种：循环链表，双向链表，带头结点的链表等





基于数组实现链表，我们也可以不适用指针，使用下标来表示下一个节点的位置，实现基于数组的链表。

### 03.带头节点的链表

带头节点的链表，好处在于所有的节点都有前驱，可以简化一部分插入和删除操作。



### 04.循环链表

循环链表是一个环状结构，链表的最后一个节点，和头节点相连。



### 05.双向链表

对比单向链表，多了一个prev节点，使得可以方便地从两个方向进行遍历

为了方便遍历的进行，指针 first 和指针last分别指向第一个节点和最后一个节点

C++中的标准list正是使用双向链表来实现

go 的标准库 container/list 也是双向链表实现的



## 2.栈

### 01.基于线性表的栈

栈是一种LIFO（LAST IN FIRST OUT 后进先出）的数据结构

基本操作

1. 创建一个栈
2. 检查栈是否为空
3. push
4. pop
5. top：返回栈顶元素

基于go slice实现的栈

```go
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

```

### 02.链式栈

链式栈不需要扩充底层数组扩容，可以充分利用存储空间。

链式栈go语言实现

```go
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

```





## 3.队列

队列是一种FIFO（FIRST IN FIRST OUT 先进先出）的数据结构

由于数组型队列浪费空间比较严重，所以不选择实现了。





### 01.链式队列

链式队列的go语言实现，基于单链表

```go
package queue

import "fmt"

// 链式队列节点
type node struct {
	element interface{}
	next    *node
}

// Queue 链式队列的结构体
type Queue struct {
	front *node
	back  *node
	len   int
}

// Init 初始化队列或者清空队列
func (q *Queue) Init() *Queue {
	q.front = nil
	q.back = nil
	q.len = 0
	return q
}

// New 获得一个初始化好的Queue
func New() *Queue {
	return new(Queue).Init()
}

// Empty 通过链式队列的元素个数len，判断队列是否为空
func (q *Queue) Empty() bool {
	if q.len == 0 {
		return true
	}
	return false
}

// Enqueue 在队列尾部添加一个元素，如果队列为空，则队列头部和尾部都会等于这个新节点
func (q *Queue) Enqueue(val interface{}) {
	if q.Empty() {
		q.back = &node{val, nil}
		q.front = q.back
		q.len++
	} else {
		tmp := q.front
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = &node{val, nil}
		q.back = tmp.next
		q.len++
	}
}

// Front 返回队列前端节点，没有判空逻辑，所以队列不为空时才能使用
func (q *Queue) Front() interface{} {
	return q.front.element
}

// Back 返回队列后端节点，没有判空逻辑，所以队列不为空时才能使用
func (q *Queue) Back() interface{} {
	return q.back.element
}

// Dequeue 删除队列前部节点，并返回值，如果队列为空，返回nil
func (q *Queue) Dequeue() interface{} {
	if q.Empty() {
		return nil
	}
	val := q.front.element
	q.front = q.front.next
	if q.len == 1 {
		q.back = nil
	}
	q.len--
	return val
}

// Display  输出队列中的元素
func (q *Queue) Display() {
	tmp := q.front
	for tmp != nil {
		fmt.Print(tmp.element, " ")
		tmp = tmp.next
	}
	fmt.Println()
}
```









## 4.二叉树和散列表

### 01.二叉树概念

首先我们把目光聚焦到 二分查找上。

对于一个有序的序列，二分查找，的效率为O(log n)，查找效率极高。

但是对于一般的链表，无法通过下标索引进行访问，定位中间元素，需要遍历该元素之前的所有元素。

对于链表这样的链式结构，我们把第一个二分查找的中间节点作为根节点，第二次两个二分查找可能的中点作为第二层节点。。。

最终就形成了一个树状结构，这种结构就可以方便地进行二分查找，被称为二分查找树。二叉查找树



子女 双亲 祖先 兄弟等容易理解的概念。。。

二叉树是树中的一种特例，即每个节点最多有两个子节点。

完全二叉树，除了最下层叶节点，每个节点都有两个子节点

非完全二叉树，除了叶节点，有的节点不足两个子节点



### 02.二叉树的数组表示

数组可以用来存储二叉树，

比如按从上到下从左到右的顺序放入数组。

适用于完全树的存储。

如果是非完全树，显然要空出很多节点，每一层的节点数分别为 2^(n-1)，由于是指数增长，浪费空间会很严重





### 03.二叉树的链表表示

为了更有效地使用空间并提供额外的灵活性，通常使用链表来实现二叉树

每个节点除了本节点的数据，还具有一个left和一个right指针，分别指向左子节点和右子节点



二叉树可以很自然地定义为递归数据结构

二叉树的三种遍历方式

LVR: 中序遍历方式 先遍历左子节点，再遍历父节点，再遍历右子节点

VLR: 前序遍历方式 先遍历父节点，再遍历左子节点，再遍历右子节点

LRV: 后序遍历方式  先遍历左子节点，再遍历右子节点，再遍历父节点



### 04.二叉查找树BST的实现

二叉查找树的性质：每个节点的值，比左子节点的值大，同时比右子节点的值小。

这种二叉树，可以方便地使用二分查找算法

基本操作

* 构造控的BST
* 判断BST是否为空
* 在BST中查找给定项
* 插入一项到BST中，并保持BST的性质
* 删除BST中的一项，并保持BST的性质
* 遍历BST，每个节点都要访问一次，多种遍历方法



无论查找，插入还是删除，都比较容易通过递归的方式实现



其中删除的操作稍微有些复杂

首先判断要删除的节点是不是头节点

因为头节点有一个特殊的地方就是没有父母节点，所以不像其他节点，需要改变父母节点的指针指向

我们先讨论删除的节点node不是头节点的情况，总共三种情况

1. 要删除的节点node是一个叶子节点

2. node有一个子女节点

3. node有两个子女节点

如果是第一种情况，很明显，我们只要把node父节点的指向node的指针设为空

如果是2，我们只需要把node父节点相应的指针指向node仅有的那个子女节点

无论node是左右节点都没关系，因为根据BST的性质，这样接上去大小都是符合的。

比如，如果node是父节点的右节点，那么node子树下面的节点必定都比node大

如果是第三种情况。

我们可以把node节点右边子树最小的节点赋值给他再执行删除，删掉那个最小的节点

也可以吧node左边子树最大的节点赋值给他，再删除那个节点

总之就是需要找这么一个节点，大小最接近原来的node，同时满足大于node的左子节点，小于node的右子节点。

也就是我上面说的那两个节点。



go 语言实现，基本上所有的查找，遍历都使用了递归操作

```go
//Package bst 二叉查找树的实现
package bst

import "fmt"

// 因为interface比较大小的时候需要确定类型，而且需要类型一致，valueType用于存储interface{}的类型方便修改
type valueType int

// BinNode 二叉查找树的节点
type BinNode struct {
	value int
	left  *BinNode
	right *BinNode
	// root用于表示节点属于哪个BST实例，删除的时候用于判断
	// root *BST
}

// BST 二叉查找树的结构体
type BST struct {
	root *BinNode
	len  int
}

// Init 二叉查找树BST的初始化，或者清空BST
func (b *BST) Init() *BST {
	b.root = nil
	b.len = 0
	return b
}

// New 返回一个初始化好的BST
func New() *BST {
	return new(BST).Init()
}

// Empty 通过判断根节点是否为nil，判断二叉树是否为空
func (b *BST) Empty() bool {
	if b.root == nil {
		return true
	}
	return false
}

// 由于不同interface{}的比较没有合适的方法，所以取消interface{} 用固定类型
// 感觉要做的话要先用type-switch判断类型，然后进行类型转换再比较，就算真的实现，性能损失也很大，不如直接用固定类型
// 还有一个问题是go 中type定义的类型，比如  type kkk int 和int不是同一个类型这样的设定。
// 比较两个interface接口的大小，过程中会先进行类型断言，再报错，如果出错，会触发panic
// 大于的情况，返回1，小于的情况返回-1，相等返回0
/*func compareInterface(v1 interface{}, v2 interface{}) int {
	v1t := v1.(type)
	v2t := v2.(type)
	if v1t != v2t {
		panic("v1 and v2 is different type")
	}
	if v1t(v1) > v2t(v2) {
		return -1
	} else if v1t(v1) < v2t(v2) {
		return 1
	} else {
		return 0
	}
}*/

// insert 插入一个节点进入二叉查找树
// 如果二叉树原本为空树，那么根节点用这个节点替换，否则插入到合适位置
// 插入需要满足二叉查找树的性质,每个节点的值比左子女大,比右子女小
// 因为要支持相同大小的元素,所以我这边实现如果一个插入节点的值大于等于节点的值,那么放右边.
// 毕竟相同元素也不会影响搜索，只是如果要搜索到全部的相同元素，需要搜一个删一个
func insert(node *BinNode, at *BinNode) {
	/*
		if compareInterface(node.value, at.value) == -1 {
			if at.left
			insert(node, at.left)
		} else {
			insert(node, at.right)
		}*/
	if at.left == nil && node.value < at.value {
		at.left = &BinNode{value: node.value, left: nil, right: nil}
		return
	}
	if at.right == nil && node.value >= at.value {
		at.right = &BinNode{value: node.value, left: nil, right: nil}
		return
	}

	if node.value < at.value {
		insert(node, at.left)
	} else {
		insert(node, at.right)
	}

}

func (b *BST) insert(node *BinNode) {
	if b.Empty() {
		b.root = node
	} else {
		insert(node, b.root)
	}
	b.len++
}

// 插入一个指定值到BST中，是对b.insert(&BinNode{value: v, left: nil, right: nil})的封装
func (b *BST) insertValue(v int) {
	b.insert(&BinNode{value: v, left: nil, right: nil})
	b.len++
}

// Add 添加一个值到BST
func (b *BST) Add(v int) {
	b.insertValue(v)
}

// 辅助进行递归操作的函数
func search(node *BinNode, v int) *BinNode {

	if node == nil || node.value == v {
		return node
	}
	if s1 := search(node.left, v); s1 != nil {
		return s1
	}
	if s2 := search(node.right, v); s2 != nil {
		return s2
	}
	return nil
}

// Search 返回值等于v的节点
// 如果没有找到，返回空
func (b *BST) Search(v int) *BinNode {
	if b.Empty() {
		return nil
	}
	return search(b.root, v)
}

// 用于递归调用函数，打印当前节点的值
// 字符串order指定遍历使用的顺序
func traverse(node *BinNode, order string) {
	switch order {
	case "LVR":
		if node != nil {
			traverse(node.left, "LVR")
			fmt.Print(node.value, " ")
			traverse(node.right, "LVR")
		}

	case "VLR":
		if node != nil {
			fmt.Print(node.value, " ")
			traverse(node.left, "VLR")
			traverse(node.right, "VLR")
		}
	case "LRV":
		if node != nil {
			traverse(node.left, "LRV")
			traverse(node.right, "LRV")
			fmt.Print(node.value, " ")
		}
	}
}

// TraverseMid  中序遍历二叉查找树
func (b *BST) TraverseMid() {
	traverse(b.root, "LVR")
	fmt.Println()
}

// TraverseFront 前序遍历二叉查找树
func (b *BST) TraverseFront() {
	traverse(b.root, "VLR")
	fmt.Println()
}

// TraverseBack 后序遍历二叉查找树
func (b *BST) TraverseBack() {
	traverse(b.root, "LRV")
	fmt.Println()
}

// graph 用于递归在指定位置输出二叉树的节点，躺着输出
// 执行流程如下
// 1.输出当前节点的右子树，整颗右子树的深度depth+1
// 2.输出depth个数的\t调整输出位置
// 3.输出当前节点的数据
// 4.输出连线，左右子树都有输出< .只有左子树输出\   只有右子树输出/
// 5.换行，再输出左子树
func graph(node *BinNode, depth int) {
	if node == nil {
		return
	}
	graph(node.right, depth+1)
	for i := 0; i < depth; i++ {
		fmt.Print("\t")
	}
	fmt.Print(node.value)
	if node.left != nil && node.right != nil {
		fmt.Print("<")
	} else if node.left != nil && node.right == nil {
		fmt.Print("\\")
	} else if node.left == nil && node.right != nil {
		fmt.Print("/")
	}
	fmt.Println()
	graph(node.left, depth+1)
}

// Graph BST图形输出
func (b *BST) Graph() {
	graph(b.root, 0)
}

func existNode(root *BinNode, node *BinNode) *BinNode {
	if root == nil || root == node {
		return root
	}
	if s1 := existNode(root.left, node); s1 != nil {
		return s1
	}
	if s2 := existNode(root.right, node); s2 != nil {
		return s2
	}
	return nil
}

// ExistNode 判断一个节点是否属于BST，如果存在返回true，不存在返回false
// 不能输入nil进行判断，没有相应的处理逻辑
func (b *BST) ExistNode(node *BinNode) bool {
	if existNode(b.root, node) != nil {
		return true
	}
	return false
}

// 不考虑根节点这种特殊情况，找到一般节点的父节点，如果找不到，返回nil
// 如果已经确定节点node属于root下的节点，那么返回空证明root=node
func findParent(root *BinNode, node *BinNode) *BinNode {
	if root == nil || root.left == node || root.right == node {
		return root
	}
	if s1 := findParent(root.left, node); s1 != nil {
		return s1
	}
	if s2 := findParent(root.right, node); s2 != nil {
		return s2
	}
	return nil
}
func remove(root *BinNode, node *BinNode) {
	// 情况1，2,左右节点至少有一个为空
	if node.left == nil || node.right == nil {
		// 情况一，node为叶节点，找到父节点，将父节点指向node的指针置为nil
		if node.left == nil && node.right == nil {
			parent := findParent(root, node)
			if parent.right == node {
				parent.right = nil
			} else {
				parent.left = nil
			}
			// 情况2，node左子节点不为空,父节点执行node的指针，指向node的左子节点
		} else if node.left != nil {
			parent := findParent(root, node)
			if parent.right == node {
				parent.right = node.left
			} else {
				parent.left = node.left
			}
			//  情况2，node右子节点不为空，父节点指向node的指针，指向node的右子节点
		} else if node.right != nil {
			parent := findParent(root, node)
			if parent.right == node {
				parent.right = node.right
			} else {
				parent.left = node.right
			}
		}
	} else {
		// 情况3 node左右子节点均不为空，找到node最左子节点的值，代替node的值，在对最左子节点，进行删除操作
		tmp := node.left
		for tmp.left != nil {
			tmp = tmp.left
		}
		node.value = tmp.value
		remove(node, tmp)
	}
}

// Remove 删除二叉树节点node
// 先把node是根节点的情况单独拿出来处理，因为根节点的异常之处在于，他没有父节点
// 分3中情况进行讨论
// 1.node是一个叶子节点
// 2.node有一个子女
// 3.node有两个子女
// 如果是1，只需把双亲节点的对应指针置为nil
// 如果是2，无论node是左节点不为空还是右节点不为空，只要把不为空的子女节点和node的父节点连接上即可
// 综上 1、2其实可以合为一种情况，即node的左右节点不全为空
// 如果是3，那么我们找到右侧节点，因为右侧都比node值要大，我们可以找出右边最小的最接近node，即不断往left走，直到下一个是nil
// 把node设为这个值，然后对这个节点执行删除操作。递归调用方法
func (b *BST) Remove(node *BinNode) {
	if b.root == node {
		// 只有一个根节点的情况,也就是没有子女节点时
		if b.len <= 1 {
			b.Init()
		} else {
			// 左右子节点均不为空
			if b.root.left != nil && b.root.right != nil {
				tmp := node.left
				for tmp.left != nil {
					tmp = tmp.left
				}
				node.value = tmp.value
				fmt.Println(tmp)
				remove(b.root, tmp)
				// 左节点不为空，右节点为空,根节点左移
			} else if b.root.left != nil {
				b.root = b.root.left
				// 右节点不为空，左节点为空，根节点右移
			} else if b.root.right != nil {
				b.root = b.root.right
			}

		}
		return
	}
	if b.ExistNode(node) {
		remove(b.root, node)
	} else {
		panic("the node given is not belong the BST")
	}
}

```











