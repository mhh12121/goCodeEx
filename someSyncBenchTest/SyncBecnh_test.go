package main

import "testing"
import "sync"

var i int

var mutex sync.RWMutex

var ch chan int = make(chan int, 1)

func add() { i += 1 }

func deferCall() {
	defer add()
}

func defer2Call() {
	defer add()
	defer add()
}

func defer10Call() {
	defer add()
	defer add()
	defer add()
	defer add()
	defer add()
	defer add()
	defer add()
	defer add()
	defer add()
	defer add()
}

func noDeferCall() {
	add()
}

func rlockedAdd() {
	mutex.RLock()
	add()
	mutex.RUnlock()
}

func chanAdd() {
	ch <- 1
	add()
	<-ch
}

func BenchmarkDefer(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		deferCall()
	}
}

func BenchmarkDefer2(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		defer2Call()
	}
}

func BenchmarkDefer10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		defer10Call()
	}
}

func BenchmarkNoDefer(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		noDeferCall()
	}
}

func BenchmarkRlockAdd(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		rlockedAdd()
	}
}

func BenchmarkChainAdd(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		chanAdd()
	}
}
