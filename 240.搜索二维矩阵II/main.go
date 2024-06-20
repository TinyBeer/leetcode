package main

import "sort"

/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
*/

func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		i := sort.SearchInts(row, target)
		if i < len(row) && row[i] == target {
			return true
		}
	}
	return false
}

// func searchMatrix(matrix [][]int, target int) bool {
// 	m, n := len(matrix), len(matrix[0])
// 	x, y := 0, n-1
// 	for x < m && y >= 0 {
// 		if matrix[x][y] == target {
// 			return true
// 		}
// 		if matrix[x][y] > target {
// 			y--
// 		} else {
// 			x++
// 		}
// 	}
// 	return false
// }
