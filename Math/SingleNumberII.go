package main

func singleNumber(nums []int) int {
	a := 0
	b := 0
	//state machine
	// 00 -> 10 -> 01 -> 00
	for _, value := range nums {
		a = (a ^ value) & ^b
		b = (b ^ value) & ^a
	}
	return a
}
