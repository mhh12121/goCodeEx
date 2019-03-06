package main

import "fmt"

func MergeSort(nums []int) {
	mSort(nums, &nums, 0, len(nums)-1)
	fmt.Println(nums)
}
func mSort(sr []int, tr1 *[]int, s int, t int) {
	//t is nums length

	tr2 := make([]int, t+1)
	if s == t {
		(*tr1)[s] = sr[s]
	} else {
		mid := (s + t) / 2
		mSort(sr, &tr2, s, mid)
		mSort(sr, &tr2, mid+1, t)
		mergeRecursive(&tr2, tr1, s, mid, t) //merge tr2[s,m] and tr2[m+1,t] to tr1[s,t]
	}

}
func mergeRecursive(sr *[]int, tr *[]int, i int, m int, n int) { //i is left, n is right, m is middle, tr is result
	k := i
	j := m + 1
	for i <= m && j <= n {
		if (*sr)[i] < (*sr)[j] {
			(*tr)[k] = (*sr)[i]
			i++
		} else {
			(*tr)[k] = (*sr)[j]
			j++
		}
		k++
	}
	if i <= m { //right part already to the end (n),but left part still not doesn't reach middle(m)
		for l := 0; l <= m-i; l++ { //directly copy rest of sr to tr
			(*tr)[k+l] = (*sr)[i+l]

		}
	}
	if j <= n { //left part already to the middle (m),but right part still not doesn't reach the end(n)
		for l := 0; l <= n-j; l++ {
			(*tr)[k+l] = (*sr)[j+l]

		}
	}

}
