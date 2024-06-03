package main

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。
*/

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	vist := make([][]bool, m)
	for i := range vist {
		vist[i] = make([]bool, n)
	}

	var explore func(int, int)
	explore = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || vist[i][j] {
			return
		}
		vist[i][j] = true
		if grid[i][j] == 1 {
			explore(i+1, j)
			explore(i-1, j)
			explore(i, j+1)
			explore(i, j-1)
		}
	}

	cnt := 0
	for i := range grid {
		for j := range grid[i] {
			if !vist[i][j] && grid[i][j] == 1 {
				cnt++
				explore(i, j)
			}
			vist[i][j] = true
		}
	}
	return cnt
}
