package main

/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
*/
func moveZeroes(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 && nums[i+1] != 0 {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		} else if nums[i] == 0 && nums[i+1] == 0 {
			tmp := findNotZero(i+1, nums)
			if tmp != len(nums)-1 {
				nums[i], nums[tmp] = nums[tmp], nums[i]
			} else { //already found the last zero
				nums[i], nums[tmp] = nums[tmp], nums[i]
				return
			}
		}
	}
	return
}
func findNotZero(index int, nums []int) int {
	for nums[index] != 0 && index < len(nums) {
		index += 1
	}
	return index
}
