package main

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。
*/

// 双指针
func maxArea(height []int) int {
	i, j, max := 0, len(height)-1, 0
	for i < j {
		if height[i] < height[j] {
			v := height[i] * (j - i)
			if max < v {
				max = v
			}
			i++
		} else {
			v := height[j] * (j - i)
			if max < v {
				max = v
			}
			j--
		}
	}
	return max
}

// 优化右边查询
// func maxArea(height []int) int {
// 	n := len(height)
// 	rhighest := make([]int, n)
// 	rp := n - 1
// 	rhighest[rp] = rp
// 	for i := 1; i < n; i++ {
// 		if height[n-i-1] > height[rp] {
// 			rp = n - i - 1
// 		}
// 		rhighest[n-i-1] = rp
// 	}
// 	max := 0
// 	pre_l := 0
// 	for i := 0; i < len(height)-1; i++ {
// 		if pre_l > height[i] {
// 			continue
// 		}
// 		pre_l = height[i]
// 		for j := i + 1; j < len(height); j++ {
// 			j = rhighest[j]
// 			v := min(height[i], height[j]) * (j - i)
// 			if max < v {
// 				max = v
// 			}
// 		}
// 	}
// 	return max
// }

// 优化左边查询
// func maxArea(height []int) int {
// 	max := 0
// 	pre_l := 0
// 	for i := 0; i < len(height)-1; i++ {
// 		if pre_l > height[i] {
// 			continue
// 		}
// 		pre_l = height[i]
// 		for j := i + 1; j < len(height); j++ {
// 			v := min(height[i], height[j]) * (j - i)
// 			if max < v {
// 				max = v
// 			}
// 		}
// 	}
// 	return max
// }

// // 暴力解法
// func maxArea(height []int) int {
// 	max := 0
// 	for i := 0; i < len(height)-1; i++ {
// 		for j := i + 1; j < len(height); j++ {
// 			v := min(height[i], height[j]) * (j - i)
// 			if max < v {
// 				max = v
// 			}
// 		}
// 	}
// 	return max
// }

// func min(x, y int) int {
// 	if x > y {
// 		return y
// 	}
// 	return x
// }
