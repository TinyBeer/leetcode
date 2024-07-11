package main

import "math"

/*
超级丑数 是一个正整数，并满足其所有质因数都出现在质数数组 primes 中。

给你一个整数 n 和一个整数数组 primes ，返回第 n 个 超级丑数 。

题目数据保证第 n 个 超级丑数 在 32-bit 带符号整数范围内。
*/

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	pointers := make([]int, len(primes))
	nums := make([]int, len(primes))
	for i := range nums {
		nums[i] = 1
	}
	for i := 1; i < len(dp); i++ {
		m := math.MaxInt64
		for j := 0; j < len(nums); j++ {
			m = min(m, nums[j])
		}
		dp[i] = m
		for j := 0; j < len(nums); j++ {
			if m == nums[j] {
				pointers[j]++
				nums[j] = dp[pointers[j]] * primes[j]
			}
		}
	}
	return dp[n]
}
