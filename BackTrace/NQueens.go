package main

/*
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.



Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement,
 where 'Q' and '.' both indicate a queen and an empty space respectively.
 Example:

Input: 4
Output: [
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.
*/
func solveNQueens(n int) [][]string {
	res := make([][]string, 0)

	visited := make([]int, n) //stands for diagonals
	row := 0                  //rowumn
	for row >= 0 {
		row++         //begin with second row
		if row == n { //overflow one rowumn,means it reached the final rowumn
			addRes(&res, &visited, n)
			row--
		} else {
			visited[row] = -1
		}
		//
		for row >= 0 {
			visited[row]++         //row ++
			if visited[row] == n { //last col
				row--
				if row == -1 { //the first row,last row
					return res
				}
			} else {
				if checkDiag(&visited, row) {
					break
				}
			}
		}
	}

	return res
}
func addRes(res *[][]string, visited *[]int, n int) {
	tmpArr := make([]string, 0)
	tmp := make([]byte, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (*visited)[i] == j { //???
				tmp = append(tmp, 'Q')
			} else {
				tmp = append(tmp, '.')
			}
		}
		tmpArr = append(tmpArr, string(tmp))
	}

	*res = append(*res, tmpArr)
	return

}

func checkDiag(visited *[]int, row int) bool {
	for i := 0; i < row; i++ {
		if (*visited)[i] == (*visited)[row] || abs((*visited)[i]-(*visited)[row]) == (row-i) {
			return false
		}
	}
	return true
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

//less time: eaiser to understand
func solveNQueens_Second(n int) [][]string {
	res := make([][]string, 0)
	colUsed := make([]bool, n)
	diag45Used := make([]bool, 2*n-1)
	diag135Used := make([]bool, 2*n-1)
	chars := make([][]byte, n)
	for i := 0; i < n; i++ {
		chars[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			chars[i][j] = '.'
		}
	}
	helper(0, &chars, &res, &colUsed, &diag45Used, &diag135Used)
	return res
}
func helper(row int, chars *[][]byte, res *[][]string, colUsed *[]bool, diag45Used *[]bool, diag135Used *[]bool) {
	length := len(*chars)
	if row == length {
		temp := make([]string, 0)
		for i := 0; i < length; i++ {
			temp = append(temp, string((*chars)[i]))
		}
		*res = append(*res, temp)
	} else {
		for i := 0; i < length; i++ {
			if (*colUsed)[i] || (*diag45Used)[i+row] || (*diag135Used)[length-1-row+i] {
				continue
			}
			(*chars)[row][i] = 'Q'
			(*colUsed)[i] = true
			(*diag45Used)[i+row] = true
			(*diag135Used)[length-1-row+i] = true
			helper(row+1, chars, res, colUsed, diag45Used, diag135Used)
			(*chars)[row][i] = '.'
			(*colUsed)[i] = false
			(*diag45Used)[i+row] = false
			(*diag135Used)[length-1-row+i] = false
		}
	}

}
