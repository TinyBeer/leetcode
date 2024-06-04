package main

import "math"

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其总和大于等于 target 的长度最小的
子数组
 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
*/

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= s {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// func minSubArrayLen(target int, nums []int) int {
// 	var sub []int
// 	ans := 0
// 	for _, n := range nums {
// 		sub = append(sub, n)
// 		target -= n
// 		if target <= 0 {
// 			if ans == 0 {
// 				ans = len(sub)
// 			} else {
// 				ans = min(ans, len(sub))
// 			}
// 			if ans == 1 {
// 				return ans
// 			}
// 		}
// 		for target < 0 && len(sub) > 0 {
// 			target += sub[0]
// 			sub = sub[1:]
// 			if target <= 0 {
// 				ans = min(ans, len(sub))
// 				if ans == 1 {
// 					return ans
// 				}
// 			}
// 		}

// 	}
// 	return ans
// }
