package main

/*
给你一个正整数 num 。如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

完全平方数 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。

不能使用任何内置的库函数，如  sqrt 。
*/

func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l < r {
		mid := (l + r) / 2
		square := mid * mid
		if square == num {
			return true
		} else if square > num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r*r == num
}
