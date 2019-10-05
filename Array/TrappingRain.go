package main

func trap(height []int) int {
	res := 0
	l := 0
	r := len(height) - 1

	for l < r && height[l] <= height[l+1] { //search from left to right: find out the tallest right wall
		l++
	}

	for l < r && height[r] <= height[r-1] { //search from right  to left: find out the tallest left wall
		r--
	}

	for l < r {
		left := height[l]  //left tallest wall
		right := height[r] //right tallest wall
		if left < right {  //???
			l++
			for l < r && left >= height[l] { //left ==> left+1,
				//if left>height[left+1 ], so it does be able to hold   left(max)-height[left]
				//if left<height[l],l will be set as tallest left wall index now
				res += left - height[l]
				l++
			}
		} else {
			r--
			for l < r && right >= height[r] {
				res += right - height[r]
				r--
			}
		}

	}
	return res
}

//water= (Math.min(height[index_left],height[index_right]) - bot[index_right])* (index_right-index_left-1)
//水的面积满足上面表达式

func trapStack(height []int) int {
	tempStack := make([]int, 0)
	i := 0
	maxWater := 0
	maxBotWater := 0
	for i < len(height) {
		if len(tempStack) == 0 || height[i] <= tempStack[len(tempStack)-1] {
			tempStack = append(tempStack, i)
			i++
		} else {
			tempBot := tempStack[len(tempStack)-1]
			tempStack = tempStack[:len(tempStack)-1]
			if len(tempStack) == 0 { //no index_left
				maxBotWater = 0
			} else {
				maxBotWater = (min(height[tempStack[len(tempStack)-1]], height[i]) - height[tempBot]) * (i - tempStack[len(tempStack)-1] - 1)

			}
			maxWater += maxBotWater
		}
	}
	return maxWater
}
