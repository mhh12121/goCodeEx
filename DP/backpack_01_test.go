package main_test

import (
	"fmt"
	"testing"
)

func Test_bp_01(t *testing.T) {
	backpack_01()
}

//make the backpack contains as most value items as possible
func backpack_01() {

	var volume = 20
	//v is volume of the backpack
	e := make([]int, 1) //every item's expense volume
	e = append(e, 5, 4, 7, 2, 6)

	w := make([]int, 1) //every item's value
	w = append(w, 12, 3, 10, 3, 6)
	var n = len(e) - 1
	result := make([][]int, n+1)
	for i := range result {
		result[i] = make([]int, volume+1)
	}

	for i := 1; i <= n; i++ { //item's quantity
		for j := 1; j <= volume; j++ { //each item's volumne
			if j < e[i] {
				result[i][j] = result[i-1][j]
			} else {
				result[i][j] = findMax(result[i-1][j], result[i-1][j-e[i]]+w[i])
			}

		}
	}
	fmt.Println(result[len(e)-1][volume])
}
func findMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}
