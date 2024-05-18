package main

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
*/

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 1 {
		return triangle[0][0]
	}
	df := make([]int, n)
	df[0] = triangle[0][0]
	var ans int
	for i := 1; i < n; i++ {
		for j := i; j >= 0; j-- {
			if i == j {
				df[j] = df[j-1] + triangle[i][j]
			} else if j == 0 {
				df[j] = df[0] + triangle[i][0]
			} else {
				df[j] = min(df[j], df[j-1]) + triangle[i][j]
			}
			if i == n-1 && j == n-1 {
				ans = df[j]
			} else {
				ans = min(ans, df[j])
			}
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
