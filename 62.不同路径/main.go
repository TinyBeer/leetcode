package main

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？
*/

func uniquePaths(m int, n int) int {
	if m < n {
		m, n = n, m
	}
	ans := 1
	for i := 0; i < n-1; i++ {
		ans *= m + i
		ans /= i + 1
	}

	return ans
}
