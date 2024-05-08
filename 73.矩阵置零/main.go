package main

/*
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
*/

func setZeroes(matrix [][]int) {
	var record [][]int
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				record = append(record, []int{i, j})
			}
		}
	}
	for _, zp := range record {
		for i := range matrix {
			matrix[i][zp[1]] = 0
		}
		for i := range matrix[0] {
			matrix[zp[0]][i] = 0
		}
	}
}
