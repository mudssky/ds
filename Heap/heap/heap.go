// Package heap ,提供所有类型的堆操作只要实现heap.Interface接口
// 堆是每个节点都是子树中最小节点的树
// 堆中最小的元素时根节点 index 0
//
// 堆是一种很普遍的用于实现优先队列的方式，要创建一个优先队列，需要实现 heap interface的less方法来确定优先级
// 还有Push方法添加元素，Pop方法移除高优先级的元素
package heap

import "sort"

// Interface 类型描述了，使用这个包的方法。
// 任何类型实现了这个借口，可以当成一个最小堆使用
// 有以下约束条件
// !h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

// Init 创建堆
// Init 是一个幂等操作
// 当堆无效时也会被调用
// 时间复杂度为 O(n)当n=h.Len()
func Init(h Interface) {
	// 建堆
	n := h.Len()
	//n/2-1即最后一个节点的父节点，有一个简单的定理是子节点下标是父节点n 2*n+1和2*n+2
	// 所以说这个遍历是遍历倒数第二层及以上的节点。
	// 也就是除了最底层节点，都执行down操作，所谓down操作，就是和子节点进行比较，交换，使满足堆的性质
	// 这里是最小堆，想必是把小的元素换到上面
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Push 把元素x加入到堆中
//  The complexity is O(log n) where n = h.Len().
func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

// Pop 移除最小堆的顶点，返回那个最小元素
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	// 0即根节点位置
	h.Swap(0, n)
	down(h, 0, n)
	// 实际的数据结构执行Pop，因为根节点经过交换到了最后，Pop删除的正是最后元素
	return h.Pop()
}

// Remove 移除并且返回元素的值
// The complexity is O(log n) where n = h.Len().
func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	// 和Pop相同的设计，先把要移除的元素和最后元素交换，用up down 调整后，执行实际数据结构的Pop操作
	return h.Pop()
}

// Fix 重建堆，当某个元素改变值的时候要调用这个操作
// The complexity is O(log n) where n = h.Len().
func Fix(h Interface, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

// up算法是把新加入的元素上调到合适位置的方法
func up(h Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

// down算法是用来维持堆结构的
func down(h Interface, i0, n int) bool {
	i := i0
	for {
		// 左子节点
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		// 比较左子节点和右子节点
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		// 退出条件满足堆的性质就可以退出
		if !h.Less(j, i) {
			break
		}
		// 如果节点比子节点小，进行交换
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
