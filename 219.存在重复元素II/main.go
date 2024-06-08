package main

/*
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	tbl := make(map[int]bool, len(nums))
	for i, v := range nums {
		if ok := tbl[v]; ok {
			return true
		}
		tbl[v] = true
		if i >= k {
			tbl[nums[i-k]] = false
		}
	}
	return false
}
