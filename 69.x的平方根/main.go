package main

import "math"

/*
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
*/

// 袖珍计算器
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}

// // 二分法
// func mySqrt(x int) int {
// 	if x <= 1 {
// 		return x
// 	}
// 	l, r := 0, x-1
// 	for l < r {
// 		mid := (l + r + 1) / 2
// 		square := mid * mid
// 		if square == x {
// 			return mid
// 		} else if square > x {
// 			r = mid - 1
// 		} else {
// 			l = mid
// 		}
// 	}
// 	return r
// }
