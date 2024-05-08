package main

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/

func climbStairs(n int) int {
	var df = make([]int, n+1)
	df[0] = 1
	for i := 1; i <= n; i++ {
		df[i] += df[i-1]
		if i >= 2 {
			df[i] += df[i-2]
		}
	}
	return df[n]
}
