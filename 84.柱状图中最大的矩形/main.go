package main

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/

// 单调栈
func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}

	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]

		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	maxArea := 0
	for i := 0; i < n; i++ {
		maxArea = max(maxArea, (right[i]-left[i]-1)*heights[i])
	}
	return maxArea
}

// // 暴力解
// func largestRectangleArea(heights []int) int {
// 	n := len(heights)
// 	if n == 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return heights[0]
// 	}

// 	maxArea := 0
// 	for i := 0; i < n; i++ {
// 		minH := heights[i]
// 		for j := i; j < n; j++ {
// 			minH = min(minH, heights[j])
// 			maxArea = max(maxArea, minH*(j-i+1))
// 		}
// 	}
// 	return maxArea
// }

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
