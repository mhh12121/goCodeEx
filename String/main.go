package main

import "fmt"

func main() {
	// A := 3
	// B := 4
	// x := formString(A, B)
	// fmt.Println(x)
	// y := shortestSubString("abcdcccbddd")
	// fmt.Println(y)
	// ge := []string{"code", "cdoe", "odec", "frame", "framer"}
	// groupAnagramsEasy(ge)
	// x := "abba"
	// fmt.Println(longestPalindrome_Dynamic(x))
	x := "BBC ABCDAB ABCDABCDABDE"
	// y := "aa"
	// buddyStrings(x, y)
	target := "ABCDABD"
	fmt.Println(kmp(x, target))
}
