package main

import "fmt"

func main() {
	a := []int{1, 2, 3}

	// x := permuteUnique(a)
	// fmt.Println(x)
	// tm := make(map[int]int)
	// tMap(tm)
	// fmt.Println(tm)
	nextPermutation(a)
	fmt.Println(a)
}
