package main

import (
	"fmt"
	"testing"
)

//kind of insertion sorting
func Test_ShellSort(t *testing.T) {
	s := []int{5, 4, 8, 1, 2, 3, 4, 99, 80, 10, 99}
	temps := make([]int, len(s)+1)
	temps = append(temps, s...)
	shellSort(temps)
	fmt.Println(temps)

}
func shellSort(s []int) {
	increment := len(s)
	var j int
	for increment > 1 {
		increment = increment/3 + 1
		for i := increment + 1; i <= len(s); i++ {
			if s[i] < s[i-increment] {
				s[0] = s[i] //temporarily saved in s[0]
				for j = i - increment; j > 0 && s[0] < s[j]; j -= increment {
					s[j+increment] = s[j]

				}
				s[j+increment] = s[0]
			}
		}
	}
}
