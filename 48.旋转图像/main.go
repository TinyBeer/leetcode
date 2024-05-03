package main

/*
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
*/

func rotate(matrix [][]int) {
	swap := func(x1, y1, x2, y2 int) {
		matrix[x1][y1], matrix[x2][y2] = matrix[x2][y2], matrix[x1][y1]
	}

	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; i+j < n; j++ {
			swap(i, j, n-j-1, n-i-1)
		}
	}

	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			swap(i, j, n-i-1, j)
		}
	}
}
