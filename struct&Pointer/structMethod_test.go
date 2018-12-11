package main

import (
	"fmt"
	"testing"
)

type T struct {
	Name string
}

func (t T) M1() {
	t.Name = "name1"
}

func (t *T) M2() {
	t.Name = "name2"
}
func Test_structMethod(t *testing.T) {
	t1 := T{"t1"}

	fmt.Println("before M1:", t1.Name)
	t1.M1()
	fmt.Println("after M1", t1.Name)

	fmt.Println("before M2", t1.Name)
	t1.M2()
	fmt.Println("after M2", t1.Name)

}
