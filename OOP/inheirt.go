package main

import "fmt"

type parent struct {
}

func (p *parent) doA() {
	fmt.Println("showA")
	p.doB()
}
func (p *parent) doB() {
	fmt.Println("showB")
}

type child struct {
	parent
}

func (c *child) doB() {
	fmt.Println("child showB")
}

func main() {
	a := &child{}
	a.doA()
}
