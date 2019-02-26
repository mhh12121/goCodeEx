package main

/*
example:
{6,3,4,9,8,7,1}
此时 first ＝ 4，second = 9
从尾巴到前找到第一个大于first的数字，就是7
交换4和7，即上面的swap函数，此时序列变成6，3，7，9，8，4，1
再将second＝9以及以后的序列重新排序，让其从小到大排序，使得整体最小，即reverse一下（因为此时肯定是递减序列）
得到最终的结果：6，3，7，1，4，8，9
*/
//{1,2,3}=>{1,3,2}
//this is for finding next permutation
func nextPermutation(nums []int) {
	var i, j int
	for i = len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			tindex := i - 1
			for k := len(nums) - 1; k >= i; k-- {
				if nums[k] > nums[tindex] {
					nums[k], nums[tindex] = nums[tindex], nums[k]
					break
				}
			}
			// for j = i; j < len(nums); j++ {
			// 	if nums[j] <= nums[tindex] && nums[j] > nums[i-1] { //this is (swap)函数
			// 		tindex = j
			// 	}
			// }
			// nums[i-1], nums[tindex] = nums[tindex], nums[i-1]

			break
		}
	}
	j = len(nums) - 1
	for i < j { //reverse
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
