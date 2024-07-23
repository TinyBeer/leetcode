package main

/*
给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，
计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
*/

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i&(i-1)] + 1
	}
	return ans
}
