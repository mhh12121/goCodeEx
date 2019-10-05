package main

import "fmt"

//1. recursive version
func QuickSort(s []int, low int, high int) {

	// var pivot int
	// if low < high {
	// 	// pivot = Partition(s, low, high)
	// 	pivot := PartitionMediumThree(s, low, high)
	// 	QuickSort(s, low, pivot-1)
	// 	QuickSort(s, pivot+1, high)
	// }

	if low < high {
		pivot := Partition(s, low, high)

		QuickSort(s, low, pivot-1)
		QuickSort(s, pivot+1, high)
	}

}

func Partition(s []int, low int, high int) int {
	var pivotkey int
	pivotkey = s[low] //bottleneck
	tempHigh := high
	tempLow := low
	for low < high {
		for low < high && s[high] >= pivotkey { //to find an element<pivotkey
			high--
		}

		s[low], s[high] = s[high], s[low]
		if tempHigh == high { //reduce useless loop
			return low
		}
		for low < high && s[low] <= pivotkey {
			low++
		}
		s[low], s[high] = s[high], s[low]
		if tempLow == low {
			return low
		}
		fmt.Println(s)

	}

	return low
}

//2. medium of three
func PartitionMediumThree(s []int, low int, high int) int {
	var pivotkey int
	//--------------get medium ---------------
	m := low + (high-low)/2
	if s[low] > s[high] {
		s[low], s[high] = s[high], s[low]
	}
	if s[m] > s[high] {
		s[m], s[high] = s[high], s[m]
	}
	if s[m] > s[low] {
		s[m], s[low] = s[low], s[m]
	}
	//--------------------------------------
	pivotkey = s[low]
	for low < high {
		for low < high && s[high] >= pivotkey {
			high--
		}
		s[low], s[high] = s[high], s[low]
		for low < high && s[low] <= pivotkey {
			low++
		}
		s[low], s[high] = s[high], s[low]
	}
	return low
}

//[leetcode] find K largest
//selection method time:O(n)
// func findKthLargestQ(nums []int, k int) int {
// 	k = len(nums) - k
// 	high := len(nums) - 1
// 	low := 0
// 	for low < high {
// 		p := Partition(nums, low, high)
// 		if p < k {
// 			low = p + 1
// 		} else if p > k {
// 			high = p - 1
// 		} else {
// 			break
// 		}
// 	}
// 	return nums[k]
// }

func QuickSortIterative(nums []int) []int {
	stack := make([]int, len(nums)-1)
	top := -1
	top++
	stack[top] = 0
	top++
	stack[top] = len(nums) - 1
	for top >= 0 {
		high := stack[top]
		top--
		low := stack[top]
		top--

		p := Partition(nums, low, high)

		if p-1 > low {
			top++
			stack[top] = low
			top++
			stack[top] = p - 1
		}

		if p+1 < high {
			top++
			stack[top] = p + 1
			top++
			stack[top] = high
		}
	}
	return nums
}

func QuickSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[0] //在low上
	i := 1
	low, high := 0, len(arr)-1
	// mid := (low + high) >> 1
	for low < high {
		fmt.Println(arr)
		if arr[i] > pivot {
			arr[i], arr[high] = arr[high], arr[i]
			high--
		} else {
			arr[i], arr[low] = arr[low], arr[i]
			low++
			i++
		}

	}
	arr[low] = pivot
	QuickSort2(arr[:low])
	QuickSort2(arr[low+1:])

}
