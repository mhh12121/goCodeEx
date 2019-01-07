package main

//todo
//give operation ["p","p","p","o"] and num[1,2,3,1]
//p=push, o= pop
//after every operation, save the num's maximum*minimum value in result
//above reust:[1*1,1*2,1*3,2*3]=[1,2,3,6]
func maximumSubArr() {
	op := []string{"p", "o", "p", "p", "o"}
	num := []int{1, 2, 3, 4, 1}
	opnum := make([]int, 0)
	result := make([]int, 0)

	min := num[0]
	max := num[0]
	for i := 0; i < len(op); i++ {
		if op[i] == "p" {
			opnum = append(opnum, num[i])
			min = findMin(num[i], min)
			max = FindMax(num[i], max)
			result = append(result, max*min)
		} else if op[i] == "o" {
			if len(opnum) > 1 {
				if min == num[i] {
					// min = SearchMin()//todo
				} else if max == num[i] {
					// max = SearchMax()//todo
				} else {
					result = append(result, max*min) //do nothing
				}
			} else {

			}

		}
	}
}

func FindMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func findMin(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// func removeElement()
