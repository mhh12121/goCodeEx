package main

func main() {
	a := make([][]int, 100)
	for i := range a {
		a[i] = make([]int, 100)
	}
	DeepTraverse(a)
}

func DeepTraverse(node [][]int) {
	for i := 0; i < len(node); i++ {

	}
}
