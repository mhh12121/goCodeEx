package main

import (
	"testing"
)

type People struct {
	Name string
}

func (p *People) String() string {
	// return fmt.Sprintf("print: %v", p) //recursive call String method
	//object to String, will call this String() method again
	return ""
}

func Test_typechange(t *testing.T) {
	p := &People{}
	p.String()
}
