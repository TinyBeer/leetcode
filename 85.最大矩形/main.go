package main

/*
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，
找出只包含 1 的最大矩形，并返回其面积。
*/

func maximalRectangle(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	for j := 0; j < n; j++ { // 对于每一列，使用基于柱状图的方法
		up := make([]int, m)
		down := make([]int, m)
		stk := []int{}
		for i, l := range left {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= l[j] {
				stk = stk[:len(stk)-1]
			}
			up[i] = -1
			if len(stk) > 0 {
				up[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		stk = nil
		for i := m - 1; i >= 0; i-- {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= left[i][j] {
				stk = stk[:len(stk)-1]
			}
			down[i] = m
			if len(stk) > 0 {
				down[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		for i, l := range left {
			height := down[i] - up[i] - 1
			area := height * l[j]
			ans = max(ans, area)
		}
	}
	return
}

// // 暴力解法
// func maximalRectangle(matrix [][]byte) int {
// 	m, n := len(matrix), len(matrix[0])
// 	t := make([][]int, m)
// 	for i := 0; i < m; i++ {
// 		t[i] = make([]int, n)
// 		for j := 0; j < n; j++ {
// 			if matrix[i][j] == '1' {
// 				if j == 0 {
// 					t[i][j] = 1
// 				} else {
// 					t[i][j] = t[i][j-1] + 1
// 				}
// 			}
// 		}
// 	}
// 	ans := 0
// 	for j := 0; j < n; j++ {
// 		left, right := make([]int, m), make([]int, m)
// 		for i := 0; i < m; i++ {
// 			right[i] = m
// 		}
// 		var stack []int
// 		for i := 0; i < m; i++ {
// 			for len(stack) > 0 && t[stack[len(stack)-1]][j] >= t[i][j] {
// 				right[stack[len(stack)-1]] = i
// 				stack = stack[:len(stack)-1]
// 			}
// 			if len(stack) == 0 {
// 				left[i] = -1
// 			} else {
// 				left[i] = stack[len(stack)-1]
// 			}
// 			stack = append(stack, i)
// 		}

// 		for i := 0; i < m; i++ {
// 			ans = max(ans, (right[i]-left[i]-1)*t[i][j])
// 		}
// 	}
// 	return ans

// }

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
