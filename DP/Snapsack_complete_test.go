package main

import (
	"fmt"
	"testing"
)

//?????? cannot solve
func Test_Snapsack_complete(t *testing.T) {
	buyItems()
}
func snapsack_complete() {

}

//leetcode
//allmoeny=50
//store={20,10}//the same item quantity
//value={20,15}//correspond to store, e.g 20 items cost 20
//use all money to buy as many items as possible
//above can do 20+20=40 , return 50-40=10 //only buy store 1
//or 15+15+15=45, return 50-45=5//only buy store 2
//or 15+15+20=50, return 0//both considered
//final 3rd solution is the best
func buyItems() {
	allmoney := 50
	store := []int{20, 10, 20}
	value := []int{20, 15, 15}
	MoneySort(store, value)
	// fmt.Println(store, value)
	// returns := 0
	result := make([]int, len(store))
	j := 0
	for i := 0; i < len(store); i++ {
		for j < allmoney {
			// j = findMax(result[i], result[i-value[i-1]]+value[i]) //already choose ith item: result[j-value[i-1]]+value[i]
			j += value[i]
		}
	}
	fmt.Println(result[len(value)])
}
func MoneySort(store []int, value []int) {
	for i := 0; i < len(store); i++ {
		tmp := float64(value[i]) / float64(store[i])
		for j := i + 1; j < len(store); j++ {
			s := float64(value[j]) / float64(store[j])
			if s > tmp {
				tmp = s

				value[i], value[j] = value[j], value[i]
				store[i], store[j] = store[j], store[i]
			}
		}
	}
}

// func findMax(i, j int) int {
// 	if i > j {
// 		return i
// 	}
// 	return j
// }
