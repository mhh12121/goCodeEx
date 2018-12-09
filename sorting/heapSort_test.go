package main

import (
	"fmt"
	"testing"
)

//O(nlog2N)
//kind of selection sorting //not stable
func TestLeetCode(t *testing.T) {
	s := []int{5, 4, 8, 1, 2, 3, 4, 99, 80, 10, 99}

	d := findKthLargest(s, 5)

	t.Log(d)
	t.Log(s)
	return
}

func TestHeapSort(t *testing.T) {
	s := []int{5, 4, 8, 1, 2, 3, 4, 99, 80, 10, 99}
	length := len(s) - 1

	//1. to heapify the data first
	BuildMaxHeap(s, length)
	//2. begin to sort
	for i := length; i > 0; i-- {
		s[0], s[i] = s[i], s[0]
		// MaxHeapify(s, 0, i-1)
		MaxHeapifyI(s, 0, i-1)
	}
	fmt.Println(s)

}
func BuildMaxHeap(s []int, length int) {
	furtherestParent := length / 2 // why begin here ,cuz all of those node have children
	for i := furtherestParent; i >= 0; i-- {
		// MaxHeapify(s, i, length)
		MaxHeapifyI(s, i, length)
	}
}

//1. recursive solution
func MaxHeapify(s []int, index int, len int) {
	lc := 2*index + 1
	rc := 2*index + 2
	cMax := lc

	if lc > len {
		return
	}

	if rc <= len && s[rc] > s[lc] {
		cMax = rc
	}
	if s[cMax] > s[index] {
		s[cMax], s[index] = s[index], s[cMax]

		MaxHeapify(s, cMax, len) //if parent node is replaced by child node, need to check if new parent node fulfill  Maxheap
	}

}

//2. Iterative
func MaxHeapifyI(s []int, index int, len int) {
	temp := s[index]
	for j := 2 * index; j <= len; j *= 2 { //j=2*index means this begins with left child
		if j < len && s[j] <= s[j+1] { //check if right child>left child//maxheap
			j++ //if yes, select right child index
		}
		if temp >= s[j] { //check if parent >= maximum(leftchild,rightchild)//maxheap
			break
		}
		//if parent < maximum(leftchild,rightchild)
		s[index] = s[j]
		index = j //begin from its left child
	}
	// sort.Ints()
	s[index] = temp
}
func MinHeapifyI(s []int, index int, len int) {
	temp := s[index]
	for j := 2 * index; j <= len; j *= 2 { //j=2*index means this begins with left child
		if j < len && s[j] > s[j+1] { //check if right child>left child//maxheap
			j++ //if yes, select right child index
		}
		if temp < s[j] { //check if parent >= maximum(leftchild,rightchild)//maxheap
			break
		}
		//if parent < maximum(leftchild,rightchild)
		s[index] = s[j]
		index = j //begin from its left child
	}
	// sort.Ints()
	s[index] = temp
}

//[Leetcode]Smallest(largest) K value in Array
//!!!need to consider situation of same elements
func findKthLargest(nums []int, k int) int {
	tempnum := make([]int, 1)
	tempnum = append(tempnum, nums...)
	length := len(tempnum)
	fmt.Println(tempnum)
	beginIndex := k / 2
	for i := beginIndex; i > 0; i-- {
		MinHeapifyI(tempnum, i, k)
	}

	for i := k + 1; i < length; i++ {
		if tempnum[i] > tempnum[1] {
			// tempnum[1], tempnum[i] = tempnum[i], tempnum[0]
			tempnum[1] = tempnum[i]
			MinHeapifyI(tempnum, 1, k)
		}

	}
	fmt.Println("tempnum", tempnum)
	return tempnum[1]
}
