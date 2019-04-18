package main

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) < 2 {
		return s
	}
	res := s[0:1]
	sb := []byte(s)
	for i := 0; i < len(sb); i++ {
		tmp := findPalindrome(s, i, i)
		if len(res) < len(tmp) { //for those odd numbers
			res = tmp
		}
		tmp = findPalindrome(s, i, i+1) //!!!for those even numbers
		if len(res) < len(tmp) {
			res = tmp
		}

	}
	return res

}
func findPalindrome(s string, i int, j int) string { //begin from the center one!
	sb := []byte(s)
	for i >= 0 && j < len(s) && sb[i] == sb[j] {
		i--
		j++
	}
	return s[i+1 : j] //+1 cuz it may lower than 0
}

//Dynamic

func longestPalindrome_Dynamic(s string) string {
	res := ""
	sb := []byte(s)
	visited := make([][]bool, len(s))
	for j := 0; j < len(visited); j++ {
		visited[j] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {

		for k := i; k < len(s); k++ {
			if sb[i] == sb[k] && (k-i < 3 || visited[i+1][k-1]) {
				visited[i][k] = true
			} else {
				visited[i][k] = false
			}

			if (res == "" || k-i+1 > len(res)) && visited[i][k] { //found longer substring than res
				res = s[i : k+1]
			}

		}
	}
	return res
}
