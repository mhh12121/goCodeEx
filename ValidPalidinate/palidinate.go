package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "0p"
	t := isPalindrome(s)
	fmt.Println(t)
}
func isPalindrome(s string) bool {
	t := strings.ToLower(s)

	sbyte := []byte(t)
	fmt.Println(sbyte)
	palid := true
	j := len(sbyte) - 1
	i := 0
	for i < j {

		for (sbyte[i] < 'a' || sbyte[i] > 'z') && (sbyte[i] > '9' || sbyte[i] < '0') { //only special characters
			i++
			if i > j {
				break
			}
		}
		for sbyte[j] < 'a' || sbyte[j] > 'z' || (sbyte[j] > '9' && sbyte[j] < '0') {
			j--
			if j < i {
				break
			}
		}
		if i < j {
			if sbyte[i] != sbyte[j] {
				palid = false
			}
		}
		i++
		j--

	}
	return palid
}
