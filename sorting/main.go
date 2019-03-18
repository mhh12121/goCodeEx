package main

import "fmt"

func main() {
	s := []int{5, 4, 8, 1, 2, 3, 4, 99, 80, 10, 99}
	// // s := []int{2, 1}
	// // QuickSort(s, 0, len(s)-1)
	// // BubbleSort(s)
	// // HeapSort(s)
	// t := findKthLargestQ(s, 10)
	// // t := findKthLargestQ(s, 5)
	QuickSortIterative(s)
	fmt.Println(s)
	// fmt.Println(s)
	// y := make([]int, 50)
	// fmt.Println(len(y))
}
