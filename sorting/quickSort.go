package main

//1. recursive version
func QuickSort(s []int, low int, high int) {

	var pivot int
	// if low < high {
	// 	// pivot = Partition(s, low, high)
	// 	pivot = PartitionMediumThree(s, low, high)
	// 	QuickSort(s, low, pivot-1)
	// 	QuickSort(s, pivot+1, high)
	// }

	for low < high {
		pivot = Partition(&s, low, high)
		QuickSort(s, low, pivot-1)
		low = pivot + 1
	}

}

func Partition(s *[]int, low int, high int) int {
	var pivotkey int
	pivotkey = (*s)[low] //bottleneck
	for low < high {
		for low < high && (*s)[high] >= pivotkey {
			high--
		}
		(*s)[low], (*s)[high] = (*s)[high], (*s)[low]
		for low < high && (*s)[low] <= pivotkey {
			low++
		}
		(*s)[low], (*s)[high] = (*s)[high], (*s)[low]
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

		p := Partition(&nums, low, high)

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
