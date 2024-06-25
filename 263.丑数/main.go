package main

/*
丑数 就是只包含质因数 2、3 和 5 的正整数。

给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。
*/

func isUgly(n int) bool {
	for n >= 5 && n%5 == 0 {
		n /= 5
	}
	for n >= 3 && n%3 == 0 {
		n /= 3
	}
	for n >= 2 && n&1 == 0 {
		n /= 2
	}
	return n == 1
}
