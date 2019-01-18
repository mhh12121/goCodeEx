package main

import "sort"

/*
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

Example:

Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/
/*
Solution:The following algorithm generates the next permutation lexicographically after a given permutation. It changes the given permutation in-place.

Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
Find the largest index l greater than k such that a[k] < a[l].
Swap the value of a[k] with that of a[l].
Reverse the sequence from a[k + 1] up to and including the final element a[n].
*/
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums) //sort first so that duplicate numbers will only occur adjacently,but here actually no need

	loop(0, &res, nums)
	return res
}
func loop(index int, res *[][]int, nums []int) {
	if index == len(nums)-1 {
		tmp2 := make([]int, len(nums))
		copy(tmp2, nums)
		*res = append(*res, tmp2)
		return
	}
	// length := len(tmp) - 1
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	for i := index; i < len(tmp); i++ {
		if i != index && tmp[i] == tmp[index] {
			continue //meet duplicate
		}
		tmp[i], tmp[index] = tmp[index], tmp[i]
		loop(index+1, res, tmp)
	}
	return
}

//Use map to save used values
func loop_map(index int, res *[][]int, nums []int) {

}
