package main

import "fmt"

func main() {
	// tmp := []int{1, 6, 4}
	tmp2 := []int{0, 2, 5, 7, 8, 0, 1}
	// moveZeroes(tmp)
	// x := findMedianSortedArrays(tmp, tmp2)
	MergeSort(tmp2)
	// fmt.Println(x)
	fmt.Println(tmp2)
}
