package main

import (
	"fmt"
)

//[leetcode 152. Maximum Product Subarray]
/*
Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/
func main() {
	nums := []int{-2, 3, -4}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	max := nums[0] //localMax
	min := nums[0]
	global := nums[0] //globalMax
	if len(nums) < 2 {
		return max
	}
	for i := 1; i < len(nums); i++ {
		tmpmax := max
		max = returnMax(nums[i]*max, nums[i]*min, nums[i])
		min = returnMin(nums[i]*min, nums[i]*tmpmax, nums[i])
		if max > global {
			global = max
		}

	}
	return global
}
func returnMax(i, j, k int) int {
	if i >= j && i >= k {
		return i
	} else if j >= i && j >= k {
		return j
	}
	return k
}
func returnMin(i, j, k int) int {
	if i <= j && i <= k {
		return i
	} else if j <= i && j <= k {
		return j
	}
	return k
}
