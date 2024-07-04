package main

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的
子序列
。
*/

func lengthOfLIS(nums []int) int {
	length, n := 1, len(nums)
	if n <= 1 {
		return n
	}
	d := make([]int, n+1)
	d[length] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > d[length] {
			length++
			d[length] = nums[i]
		} else {
			l, r, p := 1, length, 0
			for l <= r {
				mid := (l + r) / 2
				if d[mid] < nums[i] {
					p = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			d[p+1] = nums[i]
		}
	}
	return length
}

// func lengthOfLIS(nums []int) int {
// 	ans := 0
// 	dp := make([]int, len(nums))
// 	for i, n := range nums {
// 		for j := i - 1; j >= 0; j-- {
// 			if nums[j] < n {
// 				dp[i] = max(dp[i], dp[j]+1)
// 				if nums[i] == n-1 || dp[j] == j+1 {
// 					break
// 				}
// 			}
// 		}
// 		if dp[i] == 0 {
// 			dp[i] = 1
// 		}
// 		ans = max(ans, dp[i])
// 	}
// 	return ans
// }
