package main

import "sort"

/*
给你一个由 无重复 正整数组成的集合 nums ，请你找出并返回其中最大的整除子集 answer ，子集中每一元素对 (answer[i], answer[j]) 都应当满足：
answer[i] % answer[j] == 0 ，或
answer[j] % answer[i] == 0
如果存在多个有效解子集，返回其中任何一个均可。
*/

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	maxSize, maxVal := 1, 1
	for i := 1; i < n; i++ {
		for j, v := range nums[:i] {
			if nums[i]%v == 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxSize {
			maxSize = dp[i]
			maxVal = nums[i]
		}
	}
	if maxSize == 1 {
		return []int{nums[0]}
	}
	res := make([]int, 0, maxSize)
	for i := n - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return res
}
