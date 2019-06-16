package main

import "fmt"

func reOrderArr(array []int) {
	var even = -1
	length := len(array)
	for i := 0; i < length; i++ {
		if array[i]%2 != 0 {
			if even >= 0 { //already found even number before this odd
				tmp := array[i]
				for j := i; j >= even; j-- {
					array[j] = array[j-1]
				}
				array[even] = tmp
				even++
			}

		} else {
			if even < 0 {
				even = i
			}
		}
	}
	fmt.Println(array)
}
