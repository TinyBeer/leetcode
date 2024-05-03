package main

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。
你可以 按任意顺序 返回答案。
*/

// 暴力
func permute(nums []int) [][]int {
	var res [][]int
	var helper []int
	var df func(int)
	df = func(mask int) {
		for i := 0; i < len(nums); i++ {
			if mask&(1<<i) == 0 {
				helper = append(helper, nums[i])
				if len(helper) == len(nums) {
					res = append(res, append([]int(nil), helper...))
				} else {
					df(mask + (1 << i))
				}
				helper = helper[:len(helper)-1]
			}
		}
	}
	df(0)
	return res
}
