package main

import "math"

/*
给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。

如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，
使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。
*/

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	a, b := nums[0], math.MaxInt
	for i := 1; i < len(nums); i++ {
		if nums[i] > b {
			return true
		} else if nums[i] > a {
			b = min(b, nums[i])
		} else {
			a = min(a, nums[i])
		}
	}
	return false
}
