package main

/*
Given two strings A and B of lowercase letters, return true if and only if we can swap two letters in A so that the result equals B.



Example 1:

Input: A = "ab", B = "ba"
Output: true
Example 2:

Input: A = "ab", B = "ab"
Output: false
*/

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	arrA := []byte(A)
	arrB := []byte(B)
	// occurDup:=0

	tmp := make(map[byte]int)
	if A == B {
		for i := 0; i < len(A); i++ {

			if v, ok := tmp[A[i]]; !ok {
				tmp[A[i]] = 1
			} else {
				tmp[A[i]] = v + 1
			}

		}
		return len(tmp) < len(A)

	}

	difSlice := make([]int, 0)
	for i := 0; i < len(A); i++ {
		if arrA[i] != arrB[i] {
			difSlice = append(difSlice, i)
		}
	}
	return len(difSlice) == 2 && arrA[difSlice[0]] == arrB[difSlice[1]] && arrA[difSlice[1]] == arrB[difSlice[0]]

}
