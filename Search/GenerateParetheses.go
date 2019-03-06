package main

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
func generateParenthesis(n int) []string {

	return traverse("", n, n)
}
func traverse(symbol string, leftcount, rightcount int) []string {
	if leftcount < 0 || rightcount < 0 {
		return nil
	}
	if leftcount == 0 && rightcount == 0 {
		return []string{symbol}
	}
	if leftcount > rightcount {
		return nil
	}
	if leftcount == rightcount {
		traverse(symbol+"(", leftcount-1, rightcount)
	}
	ret := make([]string, 0)
	tmp := traverse(symbol+"(", leftcount-1, rightcount)
	ret = append(ret, tmp...)
	tmp = traverse(symbol+")", leftcount, rightcount-1)
	ret = append(ret, tmp...)
	return ret
}
