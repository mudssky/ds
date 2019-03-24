package sort

import (
	"math/rand"
	"time"
)

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
		// 如果头部大于最大的子节点，那么就不用继续交换了，这是退出条件
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

func quicksort(nums []int, first, last int) {
	/*
		numsLen := len(nums)
		if numsLen <= 1 {
			return
		}*/
	// 可以判断左右边界作为退出条件，如果first=last，说明只有一个元素
	// fmt.Println(first, last)
	if first >= last {
		return
	}
	left := first
	right := last
	// 选取左边第一个元素作为基准值
	pivot := nums[first]
	for left < right {
		// 在右边查找<=基准值的元素,最后right即保存了元素位置
		for nums[right] > pivot {
			right--
		}
		// 在左边查找>基准的元素
		for left < right && nums[left] <= pivot {
			left++
		}
		// 如果左右两个游标还没有碰头，交换两个元素
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	// 查找结束后，确定基准的位置,right查找的时候查找的是<=基准值的元素，所以最后会停在基准值位置
	// pos := right
	// 交换基准值和基准值位置所在的元素
	nums[right], nums[first] = nums[first], nums[right]
	quicksort(nums, first, right-1)
	quicksort(nums, right+1, last)
}

// QuickSort 快速排序算法,选取一个基准值，用两个游标，分别从后往前遍历小于基准值的，从前往后遍历大于基准值的，交换这两个元素。
// 如果两个游标相遇，说明整个列表已经遍历完了，并且可以得出游标的位置，再对游标左边的部分数组，和游标右边的部分数组执行一样的操作
func QuickSort(nums []int) {
	// fmt.Println(nums)
	quicksort(nums, 0, len(nums)-1)
}

// 对两个有序数组执行的归并操作,其中mid是把数组区域分成两部分的下标,左半部分数组为l到mid，右半部分是mid+1到终点r
func merge(nums []int, l, mid, r int) {
	// 因为需要修改原数组nums的数据，复制一份left到right的数据，用于比较
	tmpNums := make([]int, r-l+1)
	for i := 0; i <= r-l; i++ {
		tmpNums[i] = nums[i+l]
	}
	// 定义左右两个游标
	left := l
	right := mid + 1
	for i := l; i <= r; i++ {
		// 如果左边的游标超过终点，说明只剩右边的数值
		if left > mid {
			nums[i] = tmpNums[right-l]
			right++
			// 如果右边的游标超过终点，说明只剩左边的数值
		} else if right > r {
			nums[i] = tmpNums[left-l]
			left++
			// 左边的数据小于右边的数据，选左边的
		} else if tmpNums[left-l] < tmpNums[right-l] {
			nums[i] = tmpNums[left-l]
			left++
			// 剩下的情况是右边的数据比较小，选右边的
		} else {
			nums[i] = tmpNums[right-l]
			right++
		}
	}
}

func mergesort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	// 递归向下
	mergesort(nums, l, mid)
	mergesort(nums, mid+1, r)
	// 归并向上
	merge(nums, l, mid, r)
}

// MergeSort 归并排序,merge操作合并两个有序数组，mergesort归并排序，将一个数组不断二分直到只有两个元素，然后不断向上merge归并
func MergeSort(nums []int) {
	mergesort(nums, 0, len(nums)-1)
}

// RadixSort 基数排序，准备10个切片分别代表0-9，先按照末尾第一位把数组里的元素放到对应的切片里。然后按从左到右0-9，从下到上（也就是切片从左到右，下标从小到大）的顺序取出
// 继续对下一位执行同一操作，那么我们拍好了后两位
// 继续执行上述操作，知道把最大的位数执行完。
func RadixSort(nums []int) {
	numsLen := len(nums)
	if numsLen <= 1 {
		return
	}
	// 1.第一步，我们要获取数组中元素的最大位数
	// 先得到数组中最大的元素
	max := nums[0]
	for i := 1; i < numsLen; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	// 得到最大元素的位数
	length := 0
	for max != 0 {
		max /= 10
		length++
	}
	// fmt.Println(length)

	// digit存储遍历的位数
	for digit := 1; digit <= length; digit++ {
		// 2. 准备10个切片数组，用于存放基数分别为0-9的元素
		radix := make([][]int, 10)
		for i := 0; i < 10; i++ {
			// 初始化10个切片
			radix[i] = []int{}
		}
		for i := 0; i < numsLen; i++ {
			// 取遍历位的数字，先除以10，让要取的位变成最后位，然后用mod 10的方法取个位
			digitNum := nums[i]
			for j := 0; j < digit-1; j++ {
				if digitNum == 0 {
					break
				}
				digitNum /= 10
			}
			// fmt.Println(digitNum)
			digitNum %= 10
			// fmt.Println(digitNum)
			// 按照这个位数上的数字放到对应的切片中
			radix[digitNum] = append(radix[digitNum], nums[i])
		}
		// 再按左到右，从下到上的顺序取出,放回原数组
		index := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < len(radix[i]); j++ {
				nums[index] = radix[i][j]
				index++
			}
		}
	}
}

// RandomInt1000 用于生成1000个随机Int的数组，随机范围100000
func RandomInt1000() (nums []int) {
	nums = make([]int, 1000)
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(100000)
	}
	return
}

// RandomInt20 用于生成20个随机Int的数组，随机范围1000
func RandomInt20() (nums []int) {
	nums = make([]int, 20)
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(1000)
	}
	return
}

// RandomInt1000s 用于生成100个随机Int的数组，随机范围100
func RandomInt1000s() (nums []int) {
	nums = make([]int, 1000)
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(50)
	}
	return
}

// RandomInt20s 用于生成20个随机Int的数组，随机范围5
func RandomInt20s() (nums []int) {
	nums = make([]int, 20)
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(5)
	}
	return
}
