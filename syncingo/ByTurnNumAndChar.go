package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup //这个东西必须加上去，保证所有goroutines能退出,或者用sleep

func PrintNumAndChar() {
	// wg.Add(2)
	numChan := make(chan struct{}, 1)
	charChan := make(chan struct{}, 1)
	numChan <- struct{}{}
	go printNum(numChan, charChan)
	go printChar(numChan, charChan)
	defer close(numChan)
	defer close(charChan)
	// select {}
	time.Sleep(time.Second * 1)
	// wg.Wait()
}

func printNum(numChan, charChan chan struct{}) {
	i := 1
	// defer wg.Done()
	for {
		if i >= 28 {
			return

		}

		<-numChan
		fmt.Print(i)
		i++
		fmt.Print(i)
		charChan <- struct{}{}
		i++

	}
}
func printChar(numChan, charChan chan struct{}) {
	i := 'A'
	// defer wg.Done()
	for {
		if i-'A' >= 26 {
			return

		}
		<-charChan

		fmt.Printf(string(i))
		i++
		fmt.Printf(string(i))
		numChan <- struct{}{}
		i++

	}

}
