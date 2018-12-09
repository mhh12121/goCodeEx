package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	// time.Sleep(1 * time.Second)
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			// time.Sleep(1 * time.Millisecond)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
