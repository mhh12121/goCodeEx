package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	a := []int{1, 2, 3}

	// x := permuteUnique(a)
	// fmt.Println(x)
	// tm := make(map[int]int)
	// tMap(tm)
	// fmt.Println(tm)
	nextPermutation(a)
	fmt.Println(a)

	x := 4
	// fmt.Println(solveNQueens(x))
	start := time.Now()
	fmt.Println(solveNQueens_Second(x))
	fmt.Println(time.Since(start))
	var str strings.Builder
	str.WriteString("a")
}
