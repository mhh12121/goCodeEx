package main

//chan.close()
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// func main() {
// 	s := time.Now()
// 	c := make(chan int, 10)
// 	go fibonacci(cap(c), c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// 	s2 := time.Since(s)
// 	fmt.Println("yo:", s2)
// }
//----------------------------
// func main() {
// 	// chantest := make(chan int)
// 	// go sum(1, chantest)
// 	// x := <-chantest
// 	// fmt.Println("chan", x)
// }
// func sum(i int, c chan int) {
// 	for j := 0; j < 2; j++ {
// 		i += j
// 	}
// 	c <- i
// }

//-------------------------
// func main() {
// c1 := make(chan int, 2)
// c2 := make(chan int, 2)
// maximum := 5
// go func() {
// 	for i := 0; i < maximum; i++ {

// 	}
// }()

// fibonacci()

// // go teller()
// Deposit(100)
// Withdraw(10)
// Withdraw(20)
// Withdraw(50)
// fmt.Println(balance)

// Withdraw(10)
// }

//another
// func fibonacci() {

// }
