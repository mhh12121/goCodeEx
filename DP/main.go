package main

import "fmt"

func main() {
	// prices := []int{1, 2, 3, 0, 2}
	// fmt.Println("best Profit:", bestTime_stateMachine(prices))

	//---------slidewindow-----------------
	// t := []int{2, 3, 4, 2, 6, 2, 5, 1}
	// res := slideWindow(t, 3)
	// fmt.Println("slide:", res)
	// t = t[:len(t)-1]
	// fmt.Println(t)
	//--------------perfectSquare-0-----------

	// fmt.Println("perfectnum:", PerfectSquare(10))
	fmt.Println("perfectnum:", PerfectSquare_dp(10))
}
