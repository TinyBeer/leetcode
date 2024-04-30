package main

import "math"

/*

给你两个整数，被除数 dividend 和除数 divisor。将两数相除，要求 不使用 乘法、除法和取余运算。

整数除法应该向零截断，也就是截去（truncate）其小数部分。例如，8.345 将被截断为 8 ，-2.7335 将被截断至 -2 。

返回被除数 dividend 除以除数 divisor 得到的 商 。

注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−231,  231 − 1] 。本题中，如果商 严格大于 231 − 1 ，则返回 231 − 1 ；如果商 严格小于 -231 ，则返回 -231 。
*/

func divide(dividend int, divisor int) int {
	if divisor == 1 || dividend == 0 {
		return dividend
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	sign := false

	if dividend > 0 {
		dividend = -dividend
		sign = !sign
	}
	if divisor > 0 {
		divisor = -divisor
		sign = !sign
	}
	ans := 0

	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1 // 注意溢出，并且不能使用除法
		if quickAdd(dividend, divisor, mid) {
			ans = mid
			if mid == math.MaxInt32 { // 注意溢出
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if sign {
		return -ans
	}
	return ans

}

// x y 为负数 z为正数   z * y >= x
func quickAdd(x int, y int, z int) bool {
	for result, add := 0, y; z > 0; z >>= 1 {
		if z&1 == 1 {
			if result < x-add {
				return false
			}
			result += add
		}

		if z != 1 {
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}

// // 暴力
// func divide(dividend int, divisor int) int {
// 	if divisor == 1 || dividend == 0 {
// 		return dividend
// 	}
// 	if dividend == math.MinInt32 && divisor == -1 {
// 		return math.MaxInt32
// 	}
// 	if divisor == math.MinInt32 {
// 		if dividend == math.MinInt32 {
// 			return 1
// 		}
// 		return 0
// 	}

// 	sign := false

// 	if dividend > 0 {
// 		dividend = -dividend
// 		sign = !sign
// 	}
// 	if divisor > 0 {
// 		divisor = -divisor
// 		sign = !sign
// 	}
// 	cnt := 0
// 	for dividend < divisor {
// 		cnt++
// 		dividend -= divisor
// 	}
// 	if sign {
// 		cnt = -cnt
// 	}

// 	return cnt
// }
