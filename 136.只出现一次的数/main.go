package main

/*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
*/

func singleNumber(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}

// func singleNumber(nums []int) int {
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}
// 	sort.Ints(nums)
// 	p := 0
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] == nums[i-1] {
// 			continue
// 		}
// 		if i-p > 1 {
// 			p = i
// 		} else {
// 			return nums[p]
// 		}
// 	}
// 	return nums[len(nums)-1]
// }
