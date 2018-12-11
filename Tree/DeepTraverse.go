package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	Node1 := &Node{Val: 2}
	Node1.Left = &Node{Val: 1}
	Node1.Right = &Node{Val: 3}
}
func deepTraverse() {

}
