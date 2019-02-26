package main

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
		index+=1
	}
	return index
}
