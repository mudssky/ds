// Package maxheap 最大堆,基于int类型的数组实现，因为接下来要实现堆排序

// 这里算是先熟悉一下堆的操作，int 类型数组的比较实用，

package maxheap

import (
	"reflect"
	"testing"
)

func TestMaxHeap_Init(t *testing.T) {
	tests := []struct {
		name string
		mh   *MaxHeap
		want *MaxHeap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mh.Init(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxHeap.Init() = %v, want %v", got, tt.want)
			}
		})
	}
}
