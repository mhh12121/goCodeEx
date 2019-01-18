package main

import (
	"fmt"
	"unicode"
)

func letterCasePermutation(S string) []string {
	res := make([]string, 0)
	tmp := []byte(S)
	if len(tmp) == 0 {
		res = append(res, "")
		return res
	}
	fmt.Println(len(tmp))
	loop(0, tmp, &res)
	return res
}

//DFS
func loop(index int, arr []byte, res *[]string) {

	if index == len(arr) {
		*res = append(*res, string(arr))
		return
	}
	if arr[index] <= '9' && arr[index] >= '0' {

		loop(index+1, arr, res)
		return
	}
	arr[index] = byte(unicode.ToUpper(rune(arr[index])))

	loop(index+1, arr, res)

	arr[index] = byte(unicode.ToLower(rune(arr[index])))

	loop(index+1, arr, res)
	// if arr[index] <= 'z' && arr[index] >= 'a' {
	// 	if index == 0 {
	// 		res = append(res, string(arr[:]))
	// 	}

	// 	tmp := arr
	// 	tmp[index] = byte(unicode.ToUpper(rune(tmp[index])))
	// 	res = append(res, string(tmp[:]))
	// 	index++
	// 	res = loop(index, arr, res)
	// 	res = loop(index, tmp, res)
	// } else if arr[index] <= 'Z' && arr[index] >= 'A' {
	// 	if index == 0 {
	// 		res = append(res, string(arr[:]))
	// 	}

	// 	tmp := arr
	// 	tmp[index] = byte(unicode.ToLower(rune(tmp[index])))
	// 	res = append(res, string(tmp[:]))
	// 	index++
	// 	res = loop(index, arr, res)
	// 	res = loop(index, tmp, res)
	// } else {
	// 	index++
	// 	res = loop(index, arr, res)
	// }
	// index++
	// return res

}
