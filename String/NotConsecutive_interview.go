package main

import (
	"fmt"
)

//given two integer A,B means the amount of 'a' is A and 'b' is B. and not allow three consecutive aaa or bbb
//e.g: A=3,B=4, abababb is accepted, but aabbbab is not accepted
func main() {
	A := 3
	B := 4
	x := formString(A, B)
	fmt.Println(x)
}
func formString(A, B int) string {
	res := make([]byte, 0)
	i := 1
	j := 1
	fmt.Println(i ^ j)

	return string(res[:])
}
