package main

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2) < len(nums1) {
		return findMedianSortedArrays(nums2, nums1)
	}
	sLen := len(nums1)
	lLen := len(nums2)
	iMin := 0
	iMax := sLen
	mid := (sLen + lLen) / 2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := mid - i
		if i < iMax && nums2[j-1] > nums1[i] { //i is too small
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			maxleft := 0
			if i == 0 {
				maxleft = nums2[j-1]
			} else if j == 0 {
				maxleft = nums1[i-1]
			} else {
				maxleft = max(nums1[i-1], nums2[j-1])
			}
			if (sLen+lLen)%2 == 1 {
				return float64(maxleft)
			}
			minright := 0
			if i == sLen {
				minright = nums2[j]
			} else if j == lLen {
				minright = nums1[i]
			} else {
				minright = min(nums2[j], nums1[i])
			}
			return (float64(maxleft) + float64(minright)) / 2
		}

	}
	return float64(0)

}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
