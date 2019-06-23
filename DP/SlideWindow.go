package main

/* given an array, e.g : {2,3,4,2,6,2,5,1} set slideWindow is 3:  {[2,3,4],2,6,2,5,1}, {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。
get maximum value in every window
then save it into an array,finally get: {4,4,6,6,6,5}
*/
func slideWindow(nums []int, size int) []int {
	res := make([]int, 0)
	if len(nums) < size || size <= 0 || len(nums) < 0 {
		return res
	}

	dq := make([]int, 0) //save nums index
	//size is the window's length
	for i := 0; i < len(nums); i++ {
		for len(dq) != 0 && nums[dq[len(dq)-1]] <= nums[i] { //found larger one!,repeatedly until empty
			dq = dq[:len(dq)-1] //pop back

		}
		if len(dq) != 0 && i-dq[0]+1 > size { // dq[0]==i-size,check if window has passed

			dq = dq[1:] //pop front
		}
		dq = append(dq, i)
		if i+1 >= size { //window begins to slide
			res = append(res, nums[dq[0]])
		}
	}

	return res
}
