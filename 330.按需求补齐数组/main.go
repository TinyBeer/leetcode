package main

/*
给定一个已排序的正整数数组 nums ，和一个正整数 n 。从 [1, n] 区间内选取任意个数字补充到 nums 中，使得 [1, n] 区间内的任何数字都可以用 nums 中某几个数字的和来表示。

请返回 满足上述要求的最少需要补充的数字个数 。
*/
func minPatches(nums []int, n int) int {
	idx, x := 0, 1
	ans := 0
	for x <= n {
		if idx < len(nums) && nums[idx] <= x {
			x += nums[idx]
			idx++
		} else {
			ans++
			x *= 2
		}
	}
	return ans
}
