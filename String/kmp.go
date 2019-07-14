package main

func next(target []byte, n []int) {
	n[0] = 0
	j := 0
	for i := 1; i < len(target); i++ {
		for j > 0 && target[i] != target[j] {
			j = n[j-1]
		}
		if target[i] == target[j] {
			j++
		}
		n[i] = j

	}
}
func next2(target []byte, n []int) {
	n[0] = -1
	k := -1
	i := 0
	for i < len(target) {
		if k == -1 || target[k] == target[i] {
			i++
			k++
			n[i] = k
		} else {
			k = n[k]
		}
	}
}
func kmp(p string, target string) bool {
	pb := []byte(p)
	targetb := []byte(target)
	nextarr := make([]int, len(targetb))
	next(targetb, nextarr)
	k := 0
	for i := 0; i < len(pb); i++ {

		if pb[i] == targetb[k] {
			if k == len(targetb)-1 {
				return true //found
			}
			k++
		} else {
			k = nextarr[k]
		}
	}

	return false
}
func kmp2(p string, target string) bool {
	if len(target) == 0 {
		return false
	}
	pb := []byte(p)
	targetb := []byte(target)
	nextarr := make([]int, len(targetb))
	next2(targetb, nextarr)
	i := 0
	k := 0
	for i < len(pb) && k < len(targetb) {

		if k == -1 || pb[i] == targetb[k] {

			k++
			i++
		} else {
			k = nextarr[k]
		}
	}
	if k == len(targetb)-1 {
		return true //found
	}
	return false
}
