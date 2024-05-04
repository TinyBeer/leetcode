package main

/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
*/

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	var ans []int
	var df func(int, int, int, int)
	df = func(x int, y int, l int, w int) {
		if l <= 0 || w <= 0 {
			return
		}
		if l == 1 {
			ans = append(ans, matrix[x][y:y+w]...)
			return
		}
		if w == 1 {
			for i := 0; i < l; i++ {
				ans = append(ans, matrix[x+i][y])
			}
			return
		}
		for i := 0; i < w; i++ {
			ans = append(ans, matrix[x][y+i])
		}
		for j := 1; j < l; j++ {
			ans = append(ans, matrix[x+j][y+w-1])
		}
		for i := 1; i < w; i++ {
			ans = append(ans, matrix[x+l-1][y+w-i-1])
		}
		for j := 1; j < l-1; j++ {
			ans = append(ans, matrix[x+l-j-1][y])
		}
		df(x+1, y+1, l-2, w-2)
	}

	df(0, 0, m, n)
	return ans
}
