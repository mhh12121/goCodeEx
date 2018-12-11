package main

import (
	"fmt"
	"testing"
)

func Test_channelSync(t *testing.T) {
	//--------------------wrong 1----------------------
	// ch := make(chan int, 1000)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		ch <- i
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		a, ok := <-ch
	// 		if !ok {
	// 			fmt.Println("close")
	// 			return
	// 		}
	// 		fmt.Println("a: ", a)
	// 	}
	// }()
	// close(ch)
	// fmt.Println("ok")
	// time.Sleep(time.Second * 100)
	//---------------------------wrong 2-------------------------
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {

		go fn(i) //a stack, so when poping will trigger fn(4) first
	}
	return <-ch
}
