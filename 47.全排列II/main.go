package main

import (
	"sort"
)

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
*/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	helper := make([]int, n)
	var res [][]int
	var df func(int, int, int)
	df = func(mask int, cur int, pos int) {
		if cur == n {
			res = append(res, append([]int(nil), helper...))
			return
		}
		if cur > 0 && nums[cur] != nums[cur-1] {
			pos = 0
		}
		for ; pos < n; pos++ {
			if mask&(1<<pos) == 0 {
				helper[pos] = nums[cur]
				df(mask+(1<<pos), cur+1, pos+1)
			}
		}
	}
	df(0, 0, 0)
	return res
}
