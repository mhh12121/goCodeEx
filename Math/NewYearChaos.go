package main

/*
It's New Year's Day and everyone's in line for the Wonderland rollercoaster ride! There are a number of people queued up, and each person wears a sticker indicating their initial position in the queue. Initial positions increment by  from  at the front of the line to  at the back.

Any person in the queue can bribe the person directly in front of them to swap positions. If two people swap positions, they still wear the same sticker denoting their original places in line. One person can bribe at most two others. For example, if and  bribes , the queue will look like this: .

Fascinated by this chaotic queue, you decide you must know the minimum number of bribes that took place to get the queue into its current state!
e.g:

Input:
2,1,5,3,4
output:
3

Input:
2,5,1,3,4
output:
Too chaotic

Explanation:
1,2,3,4,5
2 bride 1 ===> 2,1,3,4,5
5 bride 4=====> 2,1,3,5,4
5 bride 3=====> 2,1,5,3,4
So that the output is 3 times;
the last person can bride mostly twice, and he can only bride the front person

*/

func checkChaos(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]-(i+1) > 2 { //if nums[i] is more than 2 far away from its original places
			return 99999 //it is too chaotic
		}
		for j := max(0, nums[i]-2); j < i; j++ { ///todo????
			if nums[j] > nums[i] {
				count++
			}
		}
	}
	return count
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
