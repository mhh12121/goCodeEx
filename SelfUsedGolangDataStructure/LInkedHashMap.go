package main

import (
	"context"
	"fmt"
)

func InsertedOrderMap(m map[Obj]int) {
	context.Background()
	//when inserting Elements,we adjust the sequence in map
	sortBySequence := func(o1, o2 *Obj) bool { //sortBy Sequence field
		return o1.Sequence < o2.Sequence
	}
	tempObjs := make([]Obj, 0)
	for k := range m {
		tempObjs = append(tempObjs, k)
	}

	By(sortBySequence).Sort(tempObjs)
	for _, k := range tempObjs {
		fmt.Printf("key:%v,value:%v \n", k, m[k])
	}
}
