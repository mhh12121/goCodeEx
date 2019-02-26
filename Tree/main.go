package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	// Node1 := &Node{Val: 2}
	// Node1.Left = &Node{Val: 1}
	// Node1.Right = &Node{Val: 3}
	board := make([][]byte, 4)
	board[0] = []byte{'o', 'a', 'a', 'n'}
	board[1] = []byte{'e', 't', 'a', 'e'}
	board[2] = []byte{'i', 'h', 'k', 'r'}
	board[3] = []byte{'i', 'f', 'l', 'v'}
	words := []string{"oath", "pea", "eat", "rain"}
	x := findWords(board, words)
	fmt.Println("final", x)
}
