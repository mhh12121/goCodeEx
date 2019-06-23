package main

import "fmt"

func main() {
	s := []int{4, 2, 3, 1, 6, 5, 7}
	// // s := []int{2, 1}
	// QuickSort(s, 0, len(s)-1)
	change(s)
	// // BubbleSort(s)
	// // HeapSort(s)
	// t := findKthLargestQ(s, 10)
	// // t := findKthLargestQ(s, 5)
	// QuickSortIterative(s)
	fmt.Println(s)
	// fmt.Println(s)
	// y := make([]int, 50)
	// fmt.Println(len(y))
}
func change(s []int) {
	s = append(s, 1)
}
