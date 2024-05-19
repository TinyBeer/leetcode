package main

import (
	"sort"
)

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/

// 暴力
func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	sort.Ints(nums)
	cnt, ans := 1, 1
	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[i-1]
		if d == 1 {
			cnt++
			ans = max(ans, cnt)
		} else if d != 0 {
			cnt = 1
		}
	}
	return ans
}
