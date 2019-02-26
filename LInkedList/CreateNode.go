package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetNode(nums []int) *ListNode {
	res := &ListNode{Val: 0, Next: nil}
	var tmp *ListNode

	tmp = res
	// tmp.Val = nums[0]
	for i := 0; i < len(nums)-1; i++ {
		tmp.Val = nums[i]
		tmp1 := &ListNode{Val: nums[i+1]}
		tmp.Next = tmp1
		// tmp = res.Next
		tmp = tmp.Next
	}
	return res
}
func printNode(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
