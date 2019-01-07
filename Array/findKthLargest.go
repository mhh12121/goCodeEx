package main

import (
	"container/heap"
)

type tmp []int

func (arr tmp) Less(i, j int) bool {
	return arr[i] > arr[j]
}
func (arr tmp) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr tmp) Len() int {
	return len(arr)
}
func (arr *tmp) Pop() interface{} {
	old := *arr
	n := len(old)
	res := old[n-1]
	*arr = old[:n-1]
	return res
}
func (arr *tmp) Push(i interface{}) {
	*arr = append(*arr, i.(int))
}
func findKthLargest(nums []int, k int) int {
	tmpArr := make(tmp, len(nums))
	x := 0
	for i := 0; i < len(nums); i++ {
		tmpArr[i] = nums[i]
	}
	heap.Init(&tmpArr)
	for i := 0; i < len(nums); i++ {
		x = (heap.Pop(&tmpArr)).(int)
	}
	return x
}
