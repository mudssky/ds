package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSelectSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{9, 3, 5, 4, 6, 7, 8, 2, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectSort(tt.args.nums)
			fmt.Println("select:", tt.args.nums)
		})
	}
}
func TestBubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{[]int{9, 3, 5, 4, 6, 7, 8, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.nums)
			fmt.Println("bubble:", tt.args.nums)
		})
	}
}

func TestInsertSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{[]int{9, 3, 5, 4, 6, 7, 8, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertSort(tt.args.nums)
			fmt.Println("insert:", tt.args.nums)
		})
	}
}
func TestTwBubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{[]int{9, 3, 5, 4, 6, 7, 8, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TwBubbleSort(tt.args.nums)
			fmt.Println("twbubble:", tt.args.nums)
		})
	}
}

func TestTwsBubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{[]int{9, 3, 5, 4, 6, 7, 8, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TwsBubbleSort(tt.args.nums)
		})
	}
}

func TestBinInsertSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{[]int{9, 3, 5, 4, 6, 7, 8, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BinInsertSort(tt.args.nums)
		})
	}
}

func TestHeapSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{[]int{9, 3, 5, 4, 6, 7, 8, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.nums)
			fmt.Println("heapsort:", tt.args.nums)
		})
	}
}

func TestQuickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	nums1 := RandomInt1000()
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{nums1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.nums)
			fmt.Println("quicksort:", tt.args.nums)
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{[]int{9, 3, 5, 4, 6, 7, 8, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.nums)
			fmt.Println("mergesort:", tt.args.nums)
		})
	}
}

func TestRadixSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{[]int{9, 345, 524, 4454, 65445, 7212, 8212, 211, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RadixSort(tt.args.nums)
			fmt.Println("radixsort:", tt.args.nums)
		})
	}
}

func TestRandomInt1000(t *testing.T) {
	tests := []struct {
		name     string
		wantNums []int
	}{
		{"dasd", []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(RandomInt1000())
		})
	}
}

func BenchmarkSelectSort(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000()
		b.StartTimer()
		SelectSort(nums)
	}
}
func BenchmarkBubbleSort(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000()
		b.StartTimer()
		BubbleSort(nums)
	}
}

func BenchmarkInsertSort(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000()
		b.StartTimer()
		InsertSort(nums)
	}
}

func BenchmarkTwBubbleSort(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000()
		b.StartTimer()
		TwBubbleSort(nums)
	}
}

func BenchmarkTwsBubbleSort(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000()
		b.StartTimer()
		TwsBubbleSort(nums)
	}
}

func BenchmarkBinInsertSort(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000()
		b.StartTimer()
		BinInsertSort(nums)
	}
}

func BenchmarkHeapSort(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000()
		b.StartTimer()
		HeapSort(nums)
	}
}

func BenchmarkQuickSort(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000()
		b.StartTimer()
		QuickSort(nums)
	}
}
func BenchmarkMergeSort(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000()
		b.StartTimer()
		MergeSort(nums)
	}
}
func BenchmarkRadixSort(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000()
		b.StartTimer()
		RadixSort(nums)
	}
}
func BenchmarkStandardSort(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000()
		b.StartTimer()
		sort.Ints(nums)
	}
}

func BenchmarkSelectSort20(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt20()
		b.StartTimer()
		SelectSort(nums)
	}
}
func BenchmarkBubbleSort20(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt20()
		b.StartTimer()
		BubbleSort(nums)
	}
}

func BenchmarkInsertSort20(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt20()
		b.StartTimer()
		InsertSort(nums)
	}
}

func BenchmarkTwBubbleSort20(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt20()
		b.StartTimer()
		TwBubbleSort(nums)
	}
}

func BenchmarkTwsBubbleSort20(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt20()
		b.StartTimer()
		TwsBubbleSort(nums)
	}
}

func BenchmarkBinInsertSort20(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt20()
		b.StartTimer()
		BinInsertSort(nums)
	}
}

func BenchmarkHeapSort20(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt20()
		b.StartTimer()
		HeapSort(nums)
	}
}

func BenchmarkQuickSort20(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt20()
		b.StartTimer()
		QuickSort(nums)
	}
}
func BenchmarkMergeSort20(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt20()
		b.StartTimer()
		MergeSort(nums)
	}
}
func BenchmarkRadixSort20(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt20()
		b.StartTimer()
		RadixSort(nums)
	}
}
func BenchmarkStandardSort20(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt20()
		b.StartTimer()
		sort.Ints(nums)
	}
}

func BenchmarkSelectSort20s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt20s()
		b.StartTimer()
		SelectSort(nums)
	}
}
func BenchmarkBubbleSort20s(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt20s()
		b.StartTimer()
		BubbleSort(nums)
	}
}

func BenchmarkInsertSort20s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt20s()
		b.StartTimer()
		InsertSort(nums)
	}
}

func BenchmarkTwBubbleSort20s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt20s()
		b.StartTimer()
		TwBubbleSort(nums)
	}
}

func BenchmarkTwsBubbleSort20s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt20s()
		b.StartTimer()
		TwsBubbleSort(nums)
	}
}

func BenchmarkBinInsertSort20s(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt20s()
		b.StartTimer()
		BinInsertSort(nums)
	}
}

func BenchmarkHeapSort20s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt20s()
		b.StartTimer()
		HeapSort(nums)
	}
}

func BenchmarkQuickSort20s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt20s()
		b.StartTimer()
		QuickSort(nums)
	}
}
func BenchmarkMergeSort20s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt20s()
		b.StartTimer()
		MergeSort(nums)
	}
}
func BenchmarkRadixSort20s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt20s()
		b.StartTimer()
		RadixSort(nums)
	}
}
func BenchmarkStandardSort20s(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt20s()
		b.StartTimer()
		sort.Ints(nums)
	}
}

func BenchmarkSelectSort1000s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000s()
		b.StartTimer()
		SelectSort(nums)
	}
}
func BenchmarkBubbleSort1000s(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000s()
		b.StartTimer()
		BubbleSort(nums)
	}
}

func BenchmarkInsertSort1000s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000s()
		b.StartTimer()
		InsertSort(nums)
	}
}

func BenchmarkTwBubbleSort1000s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000s()
		b.StartTimer()
		TwBubbleSort(nums)
	}
}

func BenchmarkTwsBubbleSort1000s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000s()
		b.StartTimer()
		TwsBubbleSort(nums)
	}
}

func BenchmarkBinInsertSort1000s(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := RandomInt1000s()
		b.StartTimer()
		BinInsertSort(nums)
	}
}

func BenchmarkHeapSort1000s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt1000s()
		b.StartTimer()
		HeapSort(nums)
	}
}

func BenchmarkQuickSort1000s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt1000s()
		b.StartTimer()
		QuickSort(nums)
	}
}
func BenchmarkMergeSort1000s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt1000s()
		b.StartTimer()
		MergeSort(nums)
	}
}
func BenchmarkRadixSort1000s(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt1000s()
		b.StartTimer()
		RadixSort(nums)
	}
}
func BenchmarkStandardSort1000s(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		nums := RandomInt1000s()
		b.StartTimer()
		sort.Ints(nums)
	}
}
