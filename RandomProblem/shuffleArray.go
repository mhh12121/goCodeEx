package main

import (
	"math/rand"
	"time"
)

/*
Shuffle a set of numbers without duplicates.

Example:

// Init an array with set 1, 2, and 3.
int[] nums = {1,2,3};
Solution solution = new Solution(nums);

// Shuffle the array [1,2,3] and return its result. Any permutation of [1,2,3] must equally likely to be returned.
solution.shuffle();

// Resets the array back to its original configuration [1,2,3].
solution.reset();

// Returns the random shuffling of array [1,2,3].
solution.shuffle();
这道题主要是考shuffle的随机性，是不是真的做到每个牌都有机会洗到，其实就是蓄水池算法（Reservoir Sampling)

*/

type Solution struct {
	NumOrigin []int
	random    *rand.Rand
}

func Constructor(nums []int) Solution {
	sol := &Solution{
		NumOrigin: nums,
		random:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	return *sol
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.NumOrigin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	shuffled := make([]int, len(this.NumOrigin))
	perm := this.random.Perm(len(this.NumOrigin))
	for k, v := range perm {
		shuffled[k] = this.NumOrigin[v]
	}

	return shuffled
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
