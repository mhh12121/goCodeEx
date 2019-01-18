package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resListNode := &ListNode{}
	temp := 0
	tempL1Node := l1
	tempL2Node := l2

	currNode := resListNode
	for tempL1Node != nil || tempL2Node != nil {
		t1 := 0
		t2 := 0
		if tempL1Node != nil {
			t1 = tempL1Node.Val
		}
		if tempL2Node != nil {
			t2 = tempL2Node.Val
		}
		// fmt.Println("t1,t2", tempL1Node.Val, tempL2Node.Val)
		sum := t1 + t2
		currNode.Val = sum%10 + temp

		temp = sum / 10

		currNode.Next = &ListNode{}
		currNode = currNode.Next
		if tempL1Node != nil {
			fmt.Println("t1 not nil")
			tempL1Node = tempL1Node.Next
		}
		if tempL2Node != nil {
			fmt.Println("t2 not nil")
			tempL2Node = tempL2Node.Next
		}

	}

	return resListNode.Next
}

func twoSum(nums []int, target int) []int {
	// resNum := make(map[int]int)
	// result := make([]int, 0)
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] < target {
	// 		temp := target - nums[i]
	// 		if res[nums[i]] == 0 {
	// 			res[temp] = 1
	// 		} else {
	// 			result = append(result, temp, nums[i])
	// 		}

	// 	}
	// }
	// return result

	// res := make(map[int]int)
	result := make([]int, 0)
	// var temp int
	for num := range nums {
		nums = append(nums, num)
	}
	fmt.Println(nums)
	return result
}
