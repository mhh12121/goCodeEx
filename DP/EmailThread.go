package main

import "strings"

/*
Assuming we got email lists:
"a@gmail.com","b@gmail.com","a->b"
"c@gmail.com","b@gmail.com","c->b"
"b@gmail.com","a@gmail.com","b->a"

Input:above List

Output:
1,1
2,1
1,2

[num,times]: num means this is the num-th conversation, times means this is the same conversation with times-times
#todo
*/
func emailThread(s []string) [][]int {
	res := make([][]int, 0)
	visited := make(map[string]map[string]bool)
	for i := 0; i < len(s); i++ {
		x := strings.Split(s[i], ",")
		from := x[0]
		to := x[1]
		_, ok := visited[from]
		if !ok {

		}
	}
}
