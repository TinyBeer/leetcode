package main

/*
给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。

返回 你可以获得的最大乘积 。
*/

func integerBreak(n int) int {
	if n <= 4 {
		switch n {
		case 2:
			return 1
		case 3:
			return 2
		case 4:
			return 4
		}
	}

	ans := 1
	for n >= 5 {
		ans *= 3
		n -= 3
	}
	return ans * n
}
