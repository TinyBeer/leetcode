package main

/*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。
*/

func combinationSum2(candidates []int, target int) [][]int {
	tbl := make(map[int]int)
	for _, v := range candidates {
		tbl[v]++
	}
	var ans [][]int
	var tmp, nums []int
	var dfs func(int, int)
	for k := range tbl {
		nums = append(nums, k)
	}
	dfs = func(idx int, rem int) {
		if idx >= len(nums) {
			return
		}
		if rem == 0 {
			ans = append(ans, append([]int{}, tmp...))
			return
		}
		dfs(idx+1, rem)
		if rem >= nums[idx] && tbl[nums[idx]] > 0 {
			tmp = append(tmp, nums[idx])
			tbl[nums[idx]]--
			dfs(idx, rem-nums[idx])
			tmp = tmp[:len(tmp)-1]
			tbl[nums[idx]]++
		}
	}
	dfs(0, target)
	return ans
}
