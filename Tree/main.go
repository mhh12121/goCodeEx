package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	// Node1 := &Node{Val: 1}
	// Node1.Left = &Node{Val: 2}
	// Node1.Right = &Node{Val: 2}
	// Node1.Left.Left = nil
	// Node1.Right.Left = nil
	// Node1.Left.Right = &Node{Val: 3}
	// Node1.Right.Right = &Node{Val: 3}
	// fmt.Println(isSymmetricIterative(Node1))
	// board := make([][]byte, 4)
	// board[0] = []byte{'o', 'a', 'a', 'n'}
	// board[1] = []byte{'e', 't', 'a', 'e'}
	// board[2] = []byte{'i', 'h', 'k', 'r'}
	// board[3] = []byte{'i', 'f', 'l', 'v'}
	// words := []string{"oath", "pea", "eat", "rain"}
	// x := findWords(board, words)
	// fmt.Println("final", x)

	// words2 := []string{"w", "wo", "wor", "worl", "world"}
	// xy := longestWord(words2)
	// fmt.Println(xy)
	buildBlackRedTree()
}
