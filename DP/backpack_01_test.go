package main_test

import (
	"fmt"
	"testing"
)

func Test_bp_01(t *testing.T) {
	// // backpack_01()
	// var volume = 50
	// //v is volume of the backpack
	// e := make([]int, 0)       //every item's expense volume
	// e = append(e, 10, 20, 30) //, 2, 6)

	// v := make([]int, 0)         //every item's value
	// v = append(v, 60, 100, 120) //, 3, 6)
	// var n = len(e)
	// fmt.Println(backpack_01R(volume, n, e, v))
	backpack_01_1D()
}

//make the backpack contains as most value items as possible
func backpack_01() {

	var volume = 50
	//v is volume of the backpack
	e := make([]int, 1)       //every item's volume expense
	e = append(e, 10, 20, 30) //, 2, 6)

	v := make([]int, 1)         //every item's value
	v = append(v, 60, 100, 120) //, 3, 6)
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
				result[i][j] = findMax(result[i-1][j], result[i-1][j-e[i]]+v[i])
			}

		}
	}

	fmt.Println(result[n][volume])
}

//optimize the space complexity O(V)
func backpack_01_1D() {
	var volume = 50
	//v is volume of the backpack
	e := make([]int, 1)       //every item's expense volume
	e = append(e, 10, 20, 30) //, 2, 6)

	v := make([]int, 1)         //every item's value
	v = append(v, 60, 100, 120) //, 3, 6)
	var n = len(e) - 1
	result := make([]int, volume+1)

	for i := 1; i <= n; i++ { //item's quantity
		for j := volume; j >= e[i]; j-- { //each item's volumne

			result[j] = findMax(result[j], result[j-e[i]]+v[i])

		}
	}

	fmt.Println(result[volume])
}

//recursive version
func backpack_01R(volume int, n int, e []int, v []int) int {
	if n == 0 || volume == 0 {
		return 0
	}
	if e[n-1] > volume {
		return backpack_01R(volume, n-1, e, v)
	} else {
		return findMax(backpack_01R(volume, n-1, e, v), (v[n-1] + backpack_01R(volume-e[n-1], n-1, e, v)))
	}
}
func findMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}
