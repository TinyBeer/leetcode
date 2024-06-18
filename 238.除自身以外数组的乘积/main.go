package main

/*
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
*/

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	suf := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= suf
		suf *= nums[i]
	}
	return ans
}

// func productExceptSelf(nums []int) []int {
// 	n := len(nums)
// 	mull, mulr := make([]int, n), make([]int, n)
// 	mull[0] = nums[0]
// 	mulr[n-1] = nums[n-1]
// 	for i := 1; i < n; i++ {
// 		mull[i] = mull[i-1] * nums[i]
// 		mulr[n-i-1] = mulr[n-i] * nums[n-i-1]
// 	}
// 	res := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		if i == 0 {
// 			res[i] = mulr[1]
// 		} else if i == n-1 {
// 			res[i] = mull[n-2]
// 		} else {
// 			res[i] = mull[i-1] * mulr[i+1]
// 		}
// 	}
// 	return res
// }
