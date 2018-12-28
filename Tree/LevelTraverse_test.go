package main

import (
	"testing"
)

type arr struct {
	levelArr [][]int
}

func Test_levelTraverse(T *testing.T) {
	//Two solution:
	//1. how to build a queue in Golang
	//2. DFS
	var n Node
	n.Val = 1
	levelTraverse(&n)

}
func levelTraverse(root *Node) {
	var res arr
	height := 0
	level(root, res, height)
	//do reverse
}
func level(root *Node, res arr, height int) {
	if root == nil {
		return
	}

	if len(res.levelArr)-1 < height { //DFS
		t := make([]int, 0)
		t = append(t, root.Val)
		res.levelArr = append(res.levelArr, t)
	} else {
		res.levelArr[height] = append(res.levelArr[height], root.Val)
	}
	height++
	level(root.Left, res, height)
	level(root.Right, res, height)
}
