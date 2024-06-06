package main

/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
*/

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	maxRob := func(ns []int) int {
		if len(ns) == 0 {
			return 0
		}
		if len(ns) == 1 {
			return ns[0]
		}
		dp, dp2 := make([]int, len(ns)), make([]int, len(ns))
		dp[0] = ns[0]
		dp[1] = ns[1]
		dp2[1] = ns[0]
		for i := 2; i < len(ns); i++ {
			dp[i] = max(dp[i-2], dp2[i-1]) + ns[i]
			dp2[i] = max(dp[i-1], dp2[i-1])
		}
		return max(dp[len(ns)-1], dp2[len(ns)-1])
	}
	return max(maxRob(nums[1:]), nums[0]+maxRob(nums[2:len(nums)-1]))
}

// func rob(nums []int) int {
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}
// 	var count, ans int
// 	var flag bool
// 	var dfs func(x int)
// 	dfs = func(x int) {
// 		if flag && x == len(nums)-1 {
// 			return
// 		}
// 		count += nums[x]
// 		ans = max(ans, count)
// 		for i := x + 2; i < len(nums); i++ {
// 			dfs(i)
// 		}
// 		count -= nums[x]
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		if i == 0 {
// 			flag = true
// 		}
// 		dfs(i)
// 		if i == 0 {
// 			flag = false
// 		}
// 	}
// 	return ans
// }
