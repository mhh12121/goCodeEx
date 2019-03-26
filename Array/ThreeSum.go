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
}
