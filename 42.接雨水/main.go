package main

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

// 动态规划
func trap(height []int) int {
	lMax := make([]int, len(height))
	rMax := make([]int, len(height))
	lMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		if height[i] > lMax[i-1] {
			lMax[i] = height[i]
		} else {
			lMax[i] = lMax[i-1]
		}
	}
	rMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i > 0; i-- {
		if height[i] > rMax[i+1] {
			rMax[i] = height[i]
		} else {
			rMax[i] = rMax[i+1]
		}
	}
	ans := 0
	for i := 1; i < len(height)-1; i++ {
		mh := min(lMax[i-1], rMax[i+1])
		if mh > height[i] {
			ans += mh - height[i]
		}
	}
	return ans
}

// // 单调栈
// func trap(height []int) (ans int) {
//     stack := []int{}
//     for i, h := range height {
//         for len(stack) > 0 && h > height[stack[len(stack)-1]] {
//             top := stack[len(stack)-1]
//             stack = stack[:len(stack)-1]
//             if len(stack) == 0 {
//                 break
//             }
//             left := stack[len(stack)-1]
//             curWidth := i - left - 1
//             curHeight := min(height[left], h) - height[top]
//             ans += curWidth * curHeight
//         }
//         stack = append(stack, i)
//     }
//     return
// }

// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }

// // 双指针
// func trap(height []int) int {
// 	ans := 0
// 	left, right := 0, len(height)-1
// 	leftMax, rightMax := 0, 0
// 	for left < right {
// 		leftMax = max(leftMax, height[left])
// 		rightMax = max(rightMax, height[right])
// 		if height[left] < height[right] {
// 			ans += leftMax - height[left]
// 			left++
// 		} else {
// 			ans += rightMax - height[right]
// 			right--
// 		}
// 	}
// 	return ans
// }
