package main

import (
	"fmt"
)

//recover and panic
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}
func main() {
	i := 0
	defer fmt.Println(i)
	i++
	f()
	return

	//---------------panic zero!!!!----------
	// var m map[string]string
	// m["do"] = "yes"

}
func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
