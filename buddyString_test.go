package main

import (
	"fmt"
	"testing"
)

func Test_BuddyString(t *testing.T) {
	A := "ab"
	B := "ba"
	yo := buddyStrings(A, B)
	fmt.Println(yo)
}
func buddyStrings(A string, B string) bool {
	a := []byte(A)
	b := []byte(B)
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	count := 0
	tmpchange := a[0]
	change := false
	var tmp byte
	tmp = 0

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			if tmpchange != a[i] {
				tmpchange = a[i]
				change = true
			}
		}
		if a[i] != b[i] {
			count++

			if a[i] != tmp && count < 2 {
				tmp = b[i] //the first possible value
			}
			if count == 2 && a[i] == tmp {
				continue
			} else if count == 2 && a[i] != tmp {
				return false
			}
			if count > 2 {
				return false
			}

		}

	}
	if (count == 0 && change) || (!change && count != 2) {
		return false
	}

	return true

}
