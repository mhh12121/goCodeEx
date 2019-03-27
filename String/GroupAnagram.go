package main

import (
	"sort"
	"strings"
)

/*
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
*/
//Euler and Mersenne
/*Mersenne: 2^n -1
Mersenne Prime: prime that can be convert to 2^n -1 format


*/
func groupAnagrams(strs []string) [][]string {
	tmap := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		strcount := getStrCount(strs[i])
		tmap[strcount] = append(tmap[strcount], strs[i])
	}
	var res [][]string
	for _, val := range tmap {
		res = append(res, val)
	}
	return res

}
func getStrCount(s string) string {
	var resultByte [26]byte
	for _, c := range s {
		resultByte[c-'a']++
	}
	var ts strings.Builder
	for char, count := range resultByte {
		if count > 0 {
			ts.WriteByte(byte(char + 'a'))
			ts.WriteByte(byte(count + '0'))
		}
	}
	return ts.String()
}

func groupAnagramsEasy(strs []string) [][]string {

	res := make([][]string, 0)
	tmap := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		tmpbyte := []byte(strs[i])
		sort.Slice(tmpbyte, func(i int, j int) bool {
			return tmpbyte[i] < tmpbyte[j]
		})
		t, ok := tmap[string(tmpbyte)]
		if !ok {
			temp := make([]string, 1)
			temp[0] = strs[i]
			tmap[string(tmpbyte)] = temp

		} else {
			t = append(t, strs[i])
			tmap[string(tmpbyte)] = t
		}
	}
	for _, value := range tmap {
		res = append(res, value)
	}
	return res
}
