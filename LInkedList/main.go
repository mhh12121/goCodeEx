package main

func main() {
	a := []int{2, 5, 3, 4, 6, 2, 2}
	s := GetNode(a)
	// printNode(s)
	res := swapPairs(s)
	printNode(res)
}
