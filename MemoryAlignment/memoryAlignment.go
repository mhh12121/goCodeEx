package main

import (
	"fmt"
	"reflect"
)

// sequence will affect its size
type Data struct {
	b byte
	s string
	a int32
	x int64
}
type Data1 struct {
	x int64
	a int32
	b byte
}

func main() {

	var d Data
	// unsafe.Sizeof(d)
	t := reflect.TypeOf(d)
	fmt.Println(t.Size(), t.Align())
	n := t.NumField()

	for i := 0; i < n; i++ {
		field := t.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
	var d1 Data1
	t1 := reflect.TypeOf(d1)
	fmt.Println(t1.Size(), t1.Align())
	m := t1.NumField()
	for i := 0; i < m; i++ {
		field := t1.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
	//f ,_ := t.FieldByName("b")
	//fmt.Println(f.Type.FieldAlign())
}
