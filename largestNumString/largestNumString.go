package main

import (
	"bytes"
	"sort"
	"strconv"
)

/*
Given a list of non negative integers, arrange them such that they form the largest number.

Example 1:

Input: [10,2]
Output: "210"
Example 2:

Input: [3,30,34,5,9]
Output: "9534330"

*/ //todo
func main() {
	a := make([]int, 0)
	largestNumber(a)
}

func largestNumber(nums []int) string {
	// sort.Slice()
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	sort.Sort(ByLargestNumber(strs))

	var res bytes.Buffer
	for i := 0; i < len(strs); i++ {
		res.WriteString(strs[i])
	}
	strRes := res.String()
	//whats this for??? for 0???
	if strRes[0] == '0' {
		return "0"
	}
	return strRes
}

//override!!!!!like comparator<interface>
type ByLargestNumber []string

func (n ByLargestNumber) Len() int {
	return len(n)
}
func (n ByLargestNumber) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n ByLargestNumber) Less(i, j int) bool {
	strij, strji := n[i]+n[j], n[j]+n[i]
	return !compare(strij, strji)
}

func compare(a, b string) bool {
	for i := 0; i < len(a); i++ {
		switch {
		case a[i] < b[i]:
			return true
		case a[i] > b[i]:
			return false
		}
	}
	return true
}
