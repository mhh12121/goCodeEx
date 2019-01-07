package main

func main() {
	tmp := []int{3, 2, 1, 5, 6, 4}
	findKthLargest(tmp, 2)
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
	return 0
}
