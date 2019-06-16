package main

/*
To find All Kth large numbers and return;
e.g [4,5,1,6,2,7,3,8] k=4; so return 1,2,3,4
*/
func findKthAll(input []int, k int) []int {
	// res := make([]int, 0)
	/*First method Parition:
	Just let k th elements'left is all less than kth elements
	*/

	begin := 0
	end := len(input) - 1
	if end < 1 {
		return input
	}
	index := parition(&input, begin, end)
	for index != k-1 {
		if index > k-1 {
			end = index - 1
			index = parition(&input, begin, end)
		} else if index < k-1 {
			begin = index + 1
			index = parition(&input, begin, end)
		}
	}
	return input[:k]
}
func parition(input *[]int, begin int, end int) int {
	pivot := (*input)[begin]
	for begin < end {
		for begin < end && (*input)[end] >= pivot {
			end--
		}
		(*input)[end], (*input)[begin] = (*input)[begin], (*input)[end]
		for begin < end && (*input)[begin] <= pivot {
			begin++
		}
		(*input)[end], (*input)[begin] = (*input)[begin], (*input)[end]
	}
	return begin

}
