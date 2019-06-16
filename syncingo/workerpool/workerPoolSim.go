package main

import (
	"fmt"
	"sync"
)

func genJob(n int) <-chan int { //模拟生成多个Job
	jobCh := make(chan int, 100)
	// defer close(jobCh)
	go func() {
		for i := 0; i < n; i++ {
			jobCh <- i
		}
		close(jobCh)
	}()
	return jobCh

}

func workerPool(size int, jobCh <-chan int, retChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(size)
	for i := 0; i < size; i++ {
		go worker(&wg, i, jobCh, retChan)
	}

	wg.Wait()
	close(retChan)
}
func worker(wg *sync.WaitGroup, id int, jobCh <-chan int, retChan chan<- string) {
	count := 0
	for job := range jobCh {
		count++
		ret := fmt.Sprintf("Worker %d processed job: %d, it's the %d-th processed by me", id, job, count)
		retChan <- ret

	}
	wg.Done()
}
func main() {
	jobch := genJob(50)
	retCh := make(chan string, 1000)
	workerPool(3, jobch, retCh) //
	for ret := range retCh {
		fmt.Println(ret)
	}
}
