// Package minheap 最大堆,基于int类型的数组实现，因为接下来要实现堆排序
// 这里算是先熟悉一下堆的操作，int 类型数组的比较实用，
package minheap

const (
	// INITCAP 初始容量
	INITCAP = 10
)

// MinHeap 最大堆的数据结构
type MinHeap struct {
	len  int //最大堆里的元素个数
	list []int
	cap  int //底层切片容量
}

// Init 初始化最大堆(空堆)
func (mh *MinHeap) Init() *MinHeap {
	mh.len = 0
	mh.list = make([]int, INITCAP)
	mh.cap = INITCAP
	return mh
}

// New 返回一个初始化好的空堆
func New() *MinHeap {
	return new(MinHeap).Init()
}

// InitWithCopy 初始化一个mh，使用已有的切片，将执行建堆操作
func (mh *MinHeap) InitWithCopy(toHeapify []int) *MinHeap {
	mh.list = toHeapify
	mh.len = len(toHeapify)
	mh.cap = mh.len
	mh.heapify()
	return mh
}

// NewWithCopy 传入一个已有的切片，返回一个以传入切片建堆的mh对象
func NewWithCopy(toHeapify []int) *MinHeap {
	return new(MinHeap).InitWithCopy(toHeapify)
}

// Len 返回当前堆中的元素个数
func (mh *MinHeap) Len() int {
	return mh.len
}

// Empty 检查堆是否为空
func (mh *MinHeap) Empty() bool {
	if mh.len == 0 {
		return true
	}
	return false
}

// Top 返回堆顶部的元素
func (mh *MinHeap) Top() int {
	return mh.list[0]
}

func (mh *MinHeap) heapify() {
	last := mh.len
	//  将任意一个二叉树list进行建堆的heapidy算法
	// 就是对所有非叶节点执行down操作，把一个子树的父节点都变成大于两个子节点，也就满足了堆的性质
	for i := last/2 - 1; i >= 0; i-- {
		mh.down(i)
	}
}

// More 判断堆中下标index2是否比下标index1的值小
func (mh *MinHeap) Less(index1, index2 int) bool {
	if mh.list[index1] > mh.list[index2] {
		return true
	}
	return false
}

// Swap 交换堆中两个下标所指的值
func (mh *MinHeap) Swap(index1, index2 int) {
	mh.list[index1], mh.list[index2] = mh.list[index2], mh.list[index1]
}

// expandCap 扩容底部数组为两倍
func (mh *MinHeap) expandCap() {
	mh.cap *= 2
	// 创建一个新的切片，每次容量变为两倍
	newlist := make([]int, mh.cap)
	// 将旧list数据全部拷贝到新list
	copy(newlist[:mh.len], mh.list[:mh.len])

	mh.list = newlist
}

// Push 插入一个值到最大堆mh中,执行up使这个值达到合适位置，并且维持堆次序
func (mh *MinHeap) Push(v int) {
	if mh.len == mh.cap {
		mh.expandCap()
	}
	if mh.len == 0 {
		mh.list[0] = v
		mh.len++
	} else {
		mh.list[mh.len] = v
		mh.len++
		mh.up(mh.list[mh.len-1])
	}
}

// Pop 删除最大堆顶部的值，并返回，执行down维持堆次序
func (mh *MinHeap) Pop() int {
	// 删除，元素个数减少
	mh.len--
	// 正好最后一个元素的索引是mh.len-1,和它交换，顺便原来的头部也被删除
	// lastIndex := mh.len
	rootValue := mh.list[mh.len]
	mh.Swap(0, mh.len)
	// 把这个交换来的头节点降到合适的位置
	mh.down(0)
	return rootValue
}

// up操作，适用于在堆的尾部插入节点的时候，使其上浮，维持堆次序
func (mh *MinHeap) up(nodeIndex int) {
	for {
		parent := (nodeIndex - 1) / 2
		// 如果当前节点nodeIndex是根节点，或者当前节点不比父节点大，说明满足堆次序，这是退出条件
		if parent == nodeIndex || mh.Less(nodeIndex, parent) {
			break
		}
		mh.Swap(parent, nodeIndex)
		// 进行交换影响了父节点，导致不一定满足堆次序，所以继续对父节点执行up操作
		nodeIndex = parent
	}
}

// down操作，调整一个近似堆为堆，从上到下沉降根节点的过程，传入要执行down操作的节点索引，和堆的元素个数
// 也能返回一个bool，判断某个节点是否满足堆次序，如果满足，返回false，不满足，返回true
func (mh *MinHeap) down(nodeIndex0 int) bool {
	nodeIndex := nodeIndex0
	for {
		// 先假设更大的节点是左子节点
		changeChild := nodeIndex*2 + 1
		// 小于0是数据溢出的情况，如果给的节点是叶节点的情况就会走这边break
		if changeChild >= mh.len || changeChild < 0 {
			break
		}
		// 如果右节点小于左节点
		if rightChild := changeChild + 1; rightChild < mh.len && mh.Less(changeChild, rightChild) {
			// 如果右子节点更大，要改变的节点是右子节点
			changeChild = rightChild
		}
		// 判断要判断的节点是否满足堆次序
		// 如果执行down操作的节点小于最大的子节点，要执行交换操作
		if mh.Less(nodeIndex, changeChild) {
			mh.Swap(nodeIndex, changeChild)
		} else {
			// 如果已经满足堆次序，这就是退出条件
			break
		}
		// nodeIndex这个位置已经被纠正，但是它改变了和它交换的那个节点的结构，所以下一步要继续纠正交换的那个节点
		nodeIndex = changeChild

	}
	// 加上返回值是为了判断是否有序，如果一开始就是有序的，那么nodeIndex和最开始传入的不会变化，返回false
	// 如果一开始不是有序的，返回true
	return nodeIndex > nodeIndex0

}
