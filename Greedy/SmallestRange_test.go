package main

import (
	"sort"
	"testing"
)

func Test_smallestRange(t *testing.T) {

}

//910. Smallest Range II
//given array [1,3,5] and integer K(2)
// find smallest difference between largest number and smallest numbers that are applied +K or -K
//above smallest difference = 0 (array==> [3,1,3])
func SmallestRange(A []int, K int) int {
	sort.Ints(A)
	//add -K and K ====> add 0 and 2K

	min := A[0]
	max := A[len(A)-1]
	ans := max - min
	for i := 0; i < len(A); i++ {
		min = FindMin(A[0]+2*K, A[i+1]) //A[0]+2K: min of the first half; A[i+1]: min of the second half
		max = FindMax(A[i]+2*K, max)    //A[i]+2K: max of the first half; max: max of the second half
		ans = FindMin(ans, max-min)
	}
	return ans
}
func FindMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func FindMin(i, j int) int {
	if i > j {
		return j
	}
	return i
}
