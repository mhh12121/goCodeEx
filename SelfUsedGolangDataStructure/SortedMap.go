package main

import (
	"fmt"
	"sort"
)

type Obj struct {
	Title    string
	Sequence int
	Size     int
}
type By func(o1, o2 *Obj) bool

//ObjSorter joins a By function and a slice of Objs to be sorted.
type ObjSorter struct {
	objs []Obj
	by   func(o1, o2 *Obj) bool
}

func (by By) Sort(objs []Obj) {
	os := &ObjSorter{
		objs: objs,
		by:   by,
	}
	sort.Sort(os)
}

func (s *ObjSorter) Less(i, j int) bool {
	return s.by(&s.objs[i], &s.objs[j])
}
func (s *ObjSorter) Swap(i, j int) {
	s.objs[i], s.objs[j] = s.objs[j], s.objs[i]
}
func (s *ObjSorter) Len() int {
	return len(s.objs)
}

func SortedMap(m map[Obj]int) {
	sortByTitle := func(o1, o2 *Obj) bool { //sort by Ttitle field
		return o1.Title < o2.Title
	}
	// sortBySequence := func(o1, o2 *Obj) bool { //sortBy Sequence field
	// 	return o1.Sequence < o2.Sequence
	// }
	tempObjs := make([]Obj, 0)
	for k := range m {
		tempObjs = append(tempObjs, k)
	}

	By(sortByTitle).Sort(tempObjs)
	for _, k := range tempObjs {
		fmt.Printf("key:%v,value:%v \n", k, m[k])
	}

}
