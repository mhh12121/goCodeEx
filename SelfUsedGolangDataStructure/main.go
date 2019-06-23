package main

func main() {
	obj1 := &Obj{Title: "4", Sequence: 1}
	obj2 := &Obj{Title: "2", Sequence: 2}
	obj3 := &Obj{Title: "1", Sequence: 3}
	obj4 := &Obj{Title: "1"}
	objs := make([]*Obj, 0)
	objs = append(objs, obj1, obj2, obj3, obj4)
	m := make(map[Obj]int)
	m[*obj1] = 1
	m[*obj3] = 1
	m[*obj2] = 1
	m[*obj4] = 1
	InsertedOrderMap(m)
}
