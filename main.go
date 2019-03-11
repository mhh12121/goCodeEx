package main

import (
	"fmt"
)

const (
	a = 1 << iota //1
	b             //2
	c             //4
	x = iota      //3

)

// var i int

func dosomething() func(int) int {
	i := 1
	return func(delta int) int {
		i += delta
		return i
	}
}
func FibonacciClosure() func() int {
	i, j := 0, 1
	return func() int {
		i, j = j, i+j
		fmt.Println("number now", j)
		return j
	}
}
func main() {

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

	// fmt.Println(a, b, c, x)

	s := 'a'
	fmt.Println(s)
	// x := dosomething()
	y := FibonacciClosure()
	for i := 0; i < 5; i++ {
		y()
	}
	// fmt.Println(x(1), x(2), x(3))
}
