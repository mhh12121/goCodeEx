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
