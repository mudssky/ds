// list 使用双向链表实现
// 遍历一个list的方法：
// for e:=l.Front();e!=nil;e.Next(){
// 	// 对元素e执行操作
// }

package list

// Element 双向链表list的元素
type Element struct {
	// 为了简化实现，内部的list 使用了一个环结构，root既是last的元素的下一个元素，也是first元素的前一个元素
	// 存储元素的值
	Value      interface{}
	prev, next *Element
	// 元素所属的list
	list *List
}

// Next 返回节点的下一个元素或空值
func (e *Element) Next() *Element {
	// 如果p是初始化的根节点，或者e所指的list被删除，返回nil
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev 返回节点的前一个元素或者nil
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// List 双向链表list的结构体
type List struct {
	root Element //root是双向链表的头节点
	len  int
}

// Init 初始化双向链表list，或者清空list
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New 返回一个初始化好的双向链表list
func New() *List {
	return new(List).Init()
}

// Len 返回双向链表中的元素个数
func (l *List) Len() int {
	return l.len
}

// Front 返回双向链表list的第一个元素，如果list为空，返回nil
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back 返回双向链表list的最后一个元素，如果list为空返回nil
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit 懒初始化一个list,即如果还没有初始化，执行初始化
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// 在at节点后插入一个e节点，len加一，返回e
func (l *List) insert(e, at *Element) *Element {
	n := at.next
	e.next = n
	e.prev = at
	at.next = e
	n.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue是对insert(&Element{Value: v}, at)的封装
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// remove 从list中移除一个节点e，len减一，返回e节点
func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // 避免内存泄露
	e.prev = nil //避免内存泄露
	return e
}

// move  把节点e移动到at节点右边，并且返回e
func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	n := at.next
	e.next = n
	e.prev = at
	at.next = e
	n.prev = e
	return e
}

// Remove 移除一个元素，并且返回元素的value值，这个元素不能为空
func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		// 如果e.liat==l,说明l已经初始化，并插入过节点了
		// 如过l==nil 说明e是一个零节点，这里没有处理逻辑，会crash
		l.remove(e)
	}
	return e.Value
}

// PushFront 插入一个新元素在双向链表list的头部
func (l *List) PushFront(v interface{}) *Element {
	// 懒初始化：如过还未初始化过，会执行初始化
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack 插入一个新元素在双向链表list的尾部
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore 插入一个新元素e，value值为v在mark前面，并且返回元素e
// 如果mark不是l的元素，那么list不会被修改
// mark不能为空
func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

// InsertAfter 插入一个新元素e，value值为v在mark后面，并且返回元素e
// 如果mark不是l的元素，那么list不会被修改
// mark不能为空
func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

// MoveToFront 移动元素e到list l的头部
// 如果e不是l的元素，list不会被修改
// 元素e不能为空
func (l *List) MoveToFront(e *Element) {
	// 如果e不属于这个list，或者说e本身在list的头部，不执行任何修改
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

// MoveToBack 移动元素e到list l的尾部
// 如果e不是l的元素，list不会被修改
// 元素e不能为空
func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.move(e, l.root.prev)
}

// MoveBefore 移动一个元素e到mark之前
// 如果e或者mark不是l的元素，或者e==mark，list不会被修改
// e和mark不能为nil
func (l *List) MoveBefore(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter 移动一个元素e到mark之后
// 如果e或者mark不是l的元素，或者e==mark，list不会被修改
// e和mark不能为nil
func (l *List) MoveAfter(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

// PushBackList 插入一个其他list的拷贝到list l尾部
// l和other 可以相同，但是都不能为空
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList 插入一个其他list的拷贝到list l头部
// l和other 可以相同，但是都不能为空
func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}
