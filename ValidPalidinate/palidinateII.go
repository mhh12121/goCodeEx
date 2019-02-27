package main

/*
Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

Example 1:
Input: "aba"
Output: True
Example 2:
Input: "abca"
Output: True
Explanation: You could delete the character 'c'.
*/
func validPalindromeII(s string) bool {
	sb := []byte(s)
	i := 0
	j := len(sb) - 1
	for i < j {
		if sb[i] != sb[j] {
			return find(i+1, j, sb) || find(i, j-1, sb)
		}
		i++
		j--
	}
	return true
}
func find(i, j int, s []byte) bool {

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
