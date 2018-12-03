package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//1.no chan
// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	if p == nil && q == nil {
// 		return true
// 	}
// 	if p == nil || q == nil {
// 		return false
// 	}
// 	if p.Val != q.Val {
// 		return false
// 	}
// 	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
// }

//2. with channel
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var br bool
	done := make(chan bool)
	c1 := make(chan int)
	c2 := make(chan int)
	go Walker(p, c1)
	go Walker(q, c2)
	go func() {
		for i := range c1 {
			if i == <-c2 {
				br = true
			} else {
				br = false
				break
			}
		}
		done <- true
	}()
	<-done
	return br

}
func Walker(t *TreeNode, ch chan int) {
	Walk(t, ch)
	close(ch)
}

func Walk(t *TreeNode, ch chan int) {
	if t == nil {
		return
	} else if t.Left == nil {
		ch <- t.Val
		if t.Right != nil {
			Walk(t.Right, ch)
		}
	} else {
		Walk(t.Left, ch)
	}
	ch <- t.Val
	if t.Right != nil {
		Walk(t.Right, ch)
	}

}
