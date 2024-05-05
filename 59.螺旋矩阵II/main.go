package main

/*
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
*/

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	var df func(i int, j int, cur int, n int)
	df = func(i, j, cur, n int) {
		if n <= 0 {
			return
		}
		for x := 0; x < n; x++ {
			res[i][j+x] = cur
			cur++
		}
		for x := 1; x < n; x++ {
			res[i+x][j+n-1] = cur
			cur++
		}
		for x := 1; x < n; x++ {
			res[i+n-1][j+n-x-1] = cur
			cur++
		}
		for x := 1; x < n-1; x++ {
			res[i+n-x-1][j] = cur
			cur++
		}
		df(i+1, j+1, cur, n-2)
	}
	df(0, 0, 1, n)
	return res
}
