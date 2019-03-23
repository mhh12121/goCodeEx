package main

/*
a number contains 6 or 8 in any digit will be considered as lucky number:
e.g 6666,128,166,890
but if both 6 and 8 occurs occur ,this num is not a lucky number:
e.g 668,86,1086
to calcualte how many lucky numbers in the range [low,high]
assume:
0<low<high<2^64
*/
//# todo
func countLuckyNumber6and8(low int64, high int64) int64 {
	h := getdigits(high)
	// l := getdigits(low)
	count := int64(0)
	if h <= 4 {
		count = calcNormal(low, high)
		return count
	}

	return count
}

func getdigits(n int64) int {
	count := 0
	for n != 0 {
		n = n / 10
		count++
	}
	return count
}

func genMaxDigtsNumber(digit int64, max bool) int {
	count := 1
	sum := 0
	if max { //max number
		for digit != 0 {
			sum += 9 * count
			count *= 10
			digit--
		}
		return sum
	}
	for digit != 0 {
		count *= 10
		digit--
	}
	return count
}
func calcNormal(low, high int64) int64 {
	count := int64(0)
	for low <= high {
		if check(low) {
			count++
			// fmt.Println(low)
		}
		low++
	}
	return count
}
func check(n int64) bool {
	flag6 := false
	flag8 := false
	for n != 0 {
		digit := n % 10
		if digit == 6 {
			if flag8 {
				return false
			}
			flag6 = true
		} else if digit == 8 {
			if flag6 {
				return false
			}
			flag8 = true
		}
		n = n / 10
	}
	if flag6 || flag8 {
		return true
	}
	return false
}
