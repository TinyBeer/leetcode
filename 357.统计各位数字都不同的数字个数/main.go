package main

/*
给你一个整数 n ，统计并返回各位数字都不同的数字 x 的个数，其中 0 <= x < 10n 。
*/

func countNumbersWithUniqueDigits(n int) int {
	dig := []int{9, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	ans := 0
	tmp := 1
	for i := 0; i <= n && i < len(dig); i++ {
		ans += tmp
		tmp *= dig[i]
	}
	return ans
}
