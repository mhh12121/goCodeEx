package main

import "sort"

/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { //remove uplicated
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				tmp := make([]int, 3)
				tmp[0] = nums[i]
				tmp[1] = nums[left]
				tmp[2] = nums[right]
				res = append(res, tmp)
				left++
				right--
				for left < right && nums[left] == nums[left-1] { //removed left duplicated
					left++
				}
				for left < right && nums[right] == nums[right+1] { //removed right duplicated
					right--
				}
			} else if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			}
		}

	}
	return res
	//---------[0,0,0,0] failed
	// res := make([][]int, 0)
	// if len(nums) < 3 {
	// 	return res
	// }
	// sort.Ints(nums)
	// low := 0
	// high := len(nums) - 1

	// for low < high && nums[low] <= 0 {
	// 	temp := nums[low] + nums[high]
	// 	res1 := make([]int, 3)
	// 	if temp > 0 {
	// 		x := true
	// 		for i := low + 1; i < high; i++ {
	// 			if temp > -nums[i] {
	// 				high--
	// 				x = false
	// 				break
	// 			} else if temp == (-nums[i]) {
	// 				res1[0] = nums[low]
	// 				res1[1] = nums[i]
	// 				res1[2] = nums[high]
	// 				res = append(res, res1)
	// 				low++
	// 				x = false
	// 				break
	// 			}

	// 		}
	// 		if x {
	// 			low++
	// 		}

	// 	} else if temp < 0 {
	// 		x := true
	// 		for i := high - 1; i > low; i-- {
	// 			if -temp > nums[i] {
	// 				low++
	// 				x = false
	// 				break
	// 			} else if temp == (-nums[i]) {
	// 				res1[0] = nums[low]
	// 				res1[1] = nums[i]
	// 				res1[2] = nums[high]
	// 				res = append(res, res1)
	// 				high--
	// 				x = false
	// 				break
	// 			}
	// 		}
	// 		if x {
	// 			high--
	// 		}

	// 	} else if temp == 0 {
	// 		x := true
	// 		for i := high - 1; i > low; i-- {
	// 			if nums[i] == temp {
	// 				res1[0] = nums[low]
	// 				res1[1] = nums[i]
	// 				res1[2] = nums[high]
	// 				res = append(res, res1)
	// 				high--
	// 				x = false
	// 				break
	// 			}

	// 		}
	// 		if x {
	// 			high--
	// 		}

	// 	}

	// }
	// return res
}
