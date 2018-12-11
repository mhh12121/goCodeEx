package main

import (
	"testing"
)

type student struct {
	Name string
}

func assertStu(v interface{}) {
	// switch msg := v.(type) {
	// case student, *student: //still interface
	// 	fmt.Println(msg.Name)
	// }
}

func Test_asset(t *testing.T) {
	m := map[string]student{"yoyo": {"123"}} //student here is not a pointer type,
	m["yoyo"].Name = "fadsf"
}
