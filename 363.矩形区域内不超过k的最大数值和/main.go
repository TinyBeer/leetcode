package main

import (
	"math"
)

/*
给你一个 m x n 的矩阵 matrix 和一个整数 k ，找出并返回矩阵内部矩形区域的不超过 k 的最大数值和。

题目数据保证总会存在一个数值和不超过 k 的矩形区域。
*/

func maxSumSubmatrix(matrix [][]int, k int) int {
	ans := math.MinInt
	for i := range matrix {
		for j := range matrix[i] {
			if j > 0 {
				matrix[i][j] += matrix[i][j-1]
			}
		}
		for j := range matrix[i] {
			if i > 0 {
				matrix[i][j] += matrix[i-1][j]
			}

			for x := 0; x <= i; x++ {
				for y := 0; y <= j; y++ {
					sum := matrix[i][j]
					if x > 0 {
						sum -= matrix[x-1][j]
					}
					if y > 0 {
						sum -= matrix[i][y-1]
					}
					if x > 0 && y > 0 {
						sum += matrix[x-1][y-1]
					}
					if sum == k {
						return k
					} else if sum < k {
						ans = max(ans, sum)
					}
				}
			}
			if matrix[i][j] == k {
				return k
			} else if matrix[i][j] < k {
				ans = max(ans, matrix[i][j])
			}
		}
	}
	return ans
}
