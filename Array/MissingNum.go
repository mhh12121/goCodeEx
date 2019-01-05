package main

func main() {

}

//Bitwise ???? pending
func missingNum(nums []int) int {

	result := 0
	for i, v := range nums {
		result ^= i ^ v
	}
	return result ^ len(nums)

}

//Missing Positive Nums
func missingPositiveNum(nums []int) int {

}
