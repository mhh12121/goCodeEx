package main

/*
given a string, find out the length of shortest substring contains all of characters in sting
e.g:
input:
"abccbba"
output:
3
Explanation:
"abc" is the smallest string

#todo
*/

func shortestSubString(s string) int {
	maxChar := 256
	sbyte := []byte(s)
	distinctChar := 0 //how many distinct character s
	visited := make([]bool, maxChar)
	for i := 0; i < len(s); i++ {
		if visited[sbyte[i]] == false {
			visited[sbyte[i]] = true
			distinctChar++
		}
	}
	start := 0 //character
	// start_index := -1 //result start index // for return the string result,if not , can not use
	minLength := 9999999
	count := 0 //the windows width
	currCount := make([]int, maxChar)
	for j := 0; j < len(s); j++ {
		currCount[sbyte[j]]++ //count the numbers of specific character's occurence

		//if found this character already inside
		if currCount[sbyte[j]] == 1 {
			count++
		}
		//if all characters are matched?????
		if count == distinctChar {
			//minimize the window
			for currCount[sbyte[start]] > 1 { //if found this charcater occur more than once
				if currCount[sbyte[start]] > 1 {
					currCount[sbyte[start]]--
				}
				start++ //next character
			}
			lenWindow := j - start + 1
			if minLength > lenWindow {
				minLength = lenWindow
				// start_index = start
			}
		}
	}
	return minLength
}
