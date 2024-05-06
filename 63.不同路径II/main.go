package main

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。
*/

// 动态规划
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] != 0 || obstacleGrid[0][0] != 0 {
		return 0
	}

	f := make([]int, n)
	f[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[len(f)-1]
}

// //暴力递归
// func uniquePathsWithObstacles(obstacleGrid [][]int) int {
// 	ans := 0
// 	m, n := len(obstacleGrid), len(obstacleGrid[0])
// 	if obstacleGrid[m-1][n-1] != 0 {
// 		return 0
// 	}
// 	var df func(int, int)
// 	df = func(i, j int) {
// 		if i == m-1 && j == n-1 {
// 			ans++
// 			return
// 		}
// 		if obstacleGrid[i][j] != 0 {
// 			return
// 		}

// 		if i < m-1 {
// 			df(i+1, j)
// 		}
// 		if j < n-1 {
// 			df(i, j+1)
// 		}
// 	}
// 	df(0, 0)
// 	return ans

// }
