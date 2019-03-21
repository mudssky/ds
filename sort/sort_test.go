package sort

import (
	"fmt"
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
