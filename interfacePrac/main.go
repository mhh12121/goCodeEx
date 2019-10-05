package main

import (
	"fmt"
	"log"
)

func main() {
	var x []int
	var y = []int{}

	if y == nil {
		fmt.Println("abc")
	} else {
		fmt.Println("dfddf")
	}
	log.Println(y)
	log.Println(x)

	c := Work{3}
	var a A = c
	var b B = c
	log.Println(a.ShowA())

	log.Println(b.ShowB())

}

type A interface {
	ShowA() int
}
type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}
func (w Work) ShowB() int {
	return w.i + 30
}
