package main

//find num-th ugly number
/*
Write a program to find the n-th ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

Example:

Input: n = 10
Output: 12
Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.
*/
func uglyNum(num int) int {
	if num == 0 {
		return 0
	}
	base2 := 2
	base3 := 3
	base5 := 5
	res := make([]int, num)
	res[0] = 1
	index2 := 0
	index3 := 0
	index5 := 0

	for i := 1; i < num; i++ {
		res[i] = min(min(base2, base3), base5)
		if res[i] == base2 {
			index2++
			base2 = 2 * res[index2]
		}
		if res[i] == base3 {
			index3++
			base3 = 3 * res[index3]
		}
		if res[i] == base5 {
			index5++
			base5 = 5 * res[index5]
		}
	}
	return res[num-1]
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
