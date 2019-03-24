package main

import "fmt"

func quicksort(nums []int, first, last int) {
	/*
		numsLen := len(nums)
		if numsLen <= 1 {
			return
		}*/
	// 可以判断左右边界作为退出条件，如果first=last，说明只有一个元素
	fmt.Println("qs", nums)
	if first <= last {
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
	fmt.Println(left, right)
	nums[right], nums[first] = nums[first], nums[right]
	quicksort(nums, 0, right-1)
	quicksort(nums, right+1, last)
}

// QuickSort 快速排序算法
func QuickSort(nums []int) {
	quicksort(nums, 0, len(nums)-1)
}

func main() {
	nums := []int{9, 3, 5, 4, 6, 7, 8, 2, 1}
	QuickSort(nums)
	fmt.Println(nums)
}
