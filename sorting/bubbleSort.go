package main

import (
	"fmt"
)

func BubbleSort(s []int) {

	for i := 0; i < len(s); i++ {
		flag := 0
		for j := len(s) - 2; j >= i; j-- {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}
	fmt.Println(s)
}
