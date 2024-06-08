package main

/*
给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
*/

func containsDuplicate(nums []int) bool {
	tbl := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := tbl[v]; ok {
			return true
		}
		tbl[v] = struct{}{}
	}
	return false
}
