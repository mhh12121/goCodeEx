package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

const (
	a = 1 << iota //1
	b             //2
	c             //4
	x = iota      //3
)

func main() {

	// i := 0
	// defer fmt.Println(i)
	// i++
	// return
	// f()
	// fmt.Println("return normaly from 4")
	// intputnums := []int{1, 2, 3, 4, 5}

	// target := 5
	// result := twoSum(intputnums, target)
	// fmt.Println(result)

	//----two sum-----
	// l1 := &ListNode{}
	// l1.Val = 0
	// l2 := &ListNode{}
	// l2.Val = 1
	// l2.Next = &ListNode{Val: 8, Next: nil}
	// addTwoNumbers(l1, l2)
	//------------------
	// fmt.Println(a, b, c, d, e, x)
	//---------------panic zero!!!!----------
	// var m map[string]string
	// m["do"] = "yes"
	fmt.Println(a, b, c, x)
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

//recover and panic
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
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
