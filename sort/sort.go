package sort

// SelectSort 选择排序，每次选择一个元素放到正确位置
// go语言的切片是一个引用类型，或者说一个复合类型，底层还是基础数组，如果底层的数组改变了，底层数组就和传进来的参数不同了。
// 排序的过程中，底层参数没有变化，没有调用append之类，所以可以正常改变切片的值
// 时间复杂度 O(n^2),空间复杂度O(1)
func SelectSort(nums []int) {
	numsLen := len(nums)
	// 一个变量用于保存最小值的下标
	minIndex := 0
	// 数组为空的情况不会进入循环，所以不用特殊处理
	for i := 0; i < numsLen-1; i++ {
		for j := i + 1; j < numsLen; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[minIndex], nums[i] = nums[i], nums[minIndex]
		// 重置最小值下标为下一个
		minIndex = i + 1
	}
}

// BubbleSort 冒泡排序,交换不合次序的元素，直到列表中不存在这种元素为止
// 时间复杂度 O(n^2),空间复杂度O(1)
func BubbleSort(nums []int) {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		// 用一个标记判断是否完成排序，如果遍历一遍没有要交换的数据，说明已经排好序了
		isSorted := true
		for j := 0; j < numsLen-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
}

// TwBubbleSort 即 Two Way ，双向冒泡算法
// 冒泡算法可以进行改进，也就是所谓的双向冒泡排序
// 先从左到右确定一个冒泡最大值，再从右到左冒泡一个最小值
func TwBubbleSort(nums []int) {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		// 用一个标记判断是否完成排序，如果遍历一遍没有要交换的数据，说明已经排好序了
		isSorted := true
		for j := 0; j < numsLen-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSorted = false
			}
		}
		for j := numsLen - i - 2; j >= 1; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
}

// HeapSort 用自己写的最大堆实现堆排序
/*"../Heap/maxheap"
func HeapSort(nums []int) {
	numsLen := len(nums)
	mh := maxheap.NewWithCopy(nums)

	for i := 0; i < numsLen-1; i++ {
		mh.Pop()
	}
}
*/

// down算法，用于建堆，还有删除头部后的调整工作
func down(nums []int, nodeIndex int, len int) {
	for {
		// leftChild:=2*n+1
		// 先假设要交换的是左子节点
		changeChild := 2*nodeIndex + 1

		if changeChild >= len {
			break
		}
		// 如果右子节点比左子节点大，要交换的节点是右子节点
		if rightChild := changeChild + 1; rightChild < len && nums[rightChild] > nums[changeChild] {
			changeChild = rightChild
		}
		// 如果和子节点最大的一个交换后的头部，仍然大于子节点，那么就不用继续交换了，退出条件
		if nums[nodeIndex] >= nums[changeChild] {
			break
		}
		nums[changeChild], nums[nodeIndex] = nums[nodeIndex], nums[changeChild]
		nodeIndex = changeChild
	}
}

// HeapSort 从零开始实现堆排序
func HeapSort(nums []int) {
	numsLen := len(nums)
	// heapify 在数组上建堆
	for i := numsLen/2 - 1; i >= 0; i-- {
		down(nums, i, numsLen)
	}

	// 需要一个元素记录堆的大小
	heapSize := numsLen
	// 建完堆后，是最大堆，所以数组开头是最大值，只需要不断取出最大值放到数组末尾即可
	for j := 0; j < numsLen-1; j++ {
		// 交换最大值,即头节点到尾部
		nums[0], nums[numsLen-j-1] = nums[numsLen-j-1], nums[0]
		// 已经交换到尾部的元素位置已经排好，从堆中移除，
		heapSize--
		// 此时头部可能变小，需要执行down操作，使结构满足最大堆
		down(nums, 0, heapSize)
	}
}

/*
func TwBubbleSort(nums []int) {
	numsLen := len(nums)
	// 判断是否进行了交换的标志
	isSwaped:=true
	var i int
	for isSwaped{
		isSwaped = false
        //自顶而下的扫描
        for j := i;j < numsLem- i - 1;j++ {
            if nums[j] > nums[j + 1] {
                nums[j], nums[j + 1] = nums[j + 1], nums[j]
                isSwaped = true
            }
        }

        //自底而下的扫描
        for j := numsLen - i - 1;j >= i + 1;j-- {
            if nums[j] < nums[j - 1] {
                nums[j], nums[j - 1] = nums[j - 1], nums[j]
            }
            isSwaped = true
        }

        i++
	}
}
*/

// TwsBubbleSort 在双向冒泡排序的基础上还可进行优化，因为一个循环在往左冒泡小的值，一个循环在往右冒泡大的值
// 两轮循环分别确定一个最大值和最小值，所以未排序的部分就是中间的一部分，之前的算法，已经排序的部分会重复扫描
// 我们可以设置两个位置标记，记录排序到的位置，这样每次扫描可以只扫中间的部分
func TwsBubbleSort(nums []int) {
	numsLen := len(nums)
	left := 0
	right := numsLen - 1
	for i := 0; i < numsLen; i++ {
		// 用一个标记判断是否完成排序，如果遍历一遍没有要交换的数据，说明已经排好序了
		isSorted := true
		for j := left; j < numsLen-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSorted = false
			}
		}
		right--
		for j := right; j >= 1; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				isSorted = false
			}
		}
		left++
		if isSorted {
			break
		}
	}
}

// InsertSort 反复第插入新元素到已经排好序的列表之中，另插入后的列表也是有序的。
// 但是在一个列表中操作，我们插入的时候需要先挪出位置，所以把插入点，到要插入的数据的位置，全部向后移一位
// 另外从同一个数组中操作，插入的时候要确定插入的元素，这里方法来了，我们可以直接倒着遍历，往前插，这样往后挪位置的时候也方便
// 时间复杂度O(n^2) 空间复杂度O(1)
func InsertSort(nums []int) {
	numsLen := len(nums)
	// 每次插入确定一个位置
	for i := 0; i < numsLen-1; i++ {
		// 每次要插入的元素都是在已经插入完的后一位，从这位开始往前，如果能小于前面的就往前交换
		for j := i + 1; j >= 1; j-- {
			if nums[j] <= nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}

}

// BinInsertSort 折半插入排序/二分插入排序，相比于插入排序，搜索插入位置的时候，使用了二分查找算法，而不是线性查找
// 相较于普通插入排序，元素挪移的次数是一致的，只是减少了元素比较的次数。在元素个数n比较大时，效率提升才比较明显
func BinInsertSort(nums []int) {
	numsLen := len(nums)
	// 每次插入确定一个位置
	for i := 1; i < numsLen; i++ {
		// 每次要插入的元素都是在已经插入完的后一位，从这位开始往前，如果能小于前面的就往前交换
		tmp := nums[i]
		low := 0
		high := i
		for low <= high {
			mid := (low + high) / 2
			if tmp > nums[mid] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		for j := i; j > low; j-- {
			nums[j] = nums[j-1]
		}
		nums[low] = tmp

	}

}
