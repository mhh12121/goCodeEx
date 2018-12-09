package main

import (
	"fmt"
	"testing"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

// func (t *Teacher) ShowA() {
// 	fmt.Println("teacher showA")
// }
func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func Test_inheritance(t *testing.T) {
	s := Teacher{}
	s.ShowB()
}
