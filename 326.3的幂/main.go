package main

/*
给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x
*/
func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

// func isPowerOfThree(n int) bool {
// 		if n < 1 {
// 			return false
// 		}
// 		for n > 1 {
// 			if n%3 != 0 {
// 				return false
// 			}
// 			n /= 3
// 		}
// 		return true
// }
