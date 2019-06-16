package main

import "fmt"

/*
212.WordSearch
Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring.
The same letter cell may not be used more than once in a word.
Example:

Input:
board = [
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]
words = ["oath","pea","eat","rain"]

Output: ["eat","oath"]

*/
type Trie struct {
	next [26]*Trie
	word string
}

func buildTree(w []string) *Trie {
	x := &Trie{}
	for j := 0; j < len(w); j++ {
		res := x
		tmpw := []byte(w[j])
		for i := 0; i < len(tmpw); i++ {
			tmp := tmpw[i] - 'a'
			if res.next[tmp] == nil {
				res.next[tmp] = &Trie{}
				fmt.Println("tmp:", tmp)
			}
			res = res.next[tmp]

		}
		res.word = w[j]
		fmt.Println("word", res.word)
	}

	return x

}
func findWords(board [][]byte, words []string) []string {
	// t:=&Trie{}
	t := buildTree(words)
	resList := make([]string, 0)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {

			dfs(board, i, j, t, &resList)
		}
	}

	return resList
}
func dfs(board [][]byte, i, j int, p *Trie, resList *[]string) {
	c := board[i][j]
	if c == '#' || p.next[c-'a'] == nil {
		return
	}
	p = p.next[c-'a']
	if p.word != "" {
		*resList = append(*resList, p.word)
		p.word = ""
	}
	board[i][j] = '#'
	if i > 0 {
		dfs(board, i-1, j, p, resList)
	}
	if j > 0 {
		dfs(board, i, j-1, p, resList)
	}
	if i < len(board)-1 {
		dfs(board, i+1, j, p, resList)
	}
	if j < len(board[0])-1 {
		dfs(board, i, j+1, p, resList)
	}
	board[i][j] = c

}
