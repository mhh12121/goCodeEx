package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func doRandom(cmd string, finish chan int) {
	fmt.Println("133213")
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randonNum := r.Intn(100)
	sleepTime := randonNum + 100
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	if finish != nil {
		finish <- sleepTime
	}
}

func doWorkSleep(ctx context.Context, ch chan bool) {
	defer func() {
		fmt.Println("worker command finish")
		ch <- true
	}()
	sleepChan := make(chan int)
	go doRandom("yoyo1", sleepChan)
	select {
	case <-ctx.Done():
		fmt.Println("cancel return")
	case sleeptime := <-sleepChan:
		fmt.Println("sleep cancel", sleeptime)
	}
}
func doWorkContext(ctx context.Context) {
	ctxWithTime, cancelFunc := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)
	defer func() {
		fmt.Println("doworkcontext ")
		cancelFunc()
	}()
	ch := make(chan bool)
	go doWorkSleep(ctxWithTime, ch)
	select {
	case <-ctx.Done():
		fmt.Println("doworkcontext done")
	case <-ch:
		fmt.Println("doworksleep channel return")
	}
}
func wrapDo() {
	ctx := context.Background()

	ctxwithCancel, cancelFunc := context.WithCancel(ctx)
	defer func() {
		fmt.Println("main defer cancelfunc")
		cancelFunc()

	}()
	go func() {
		doRandom("main", nil)
		cancelFunc()
		fmt.Println("main sleep complete, canceling context")
	}()
	doWorkContext(ctxwithCancel)
}
