package main

/*
给定一个整数 n ，返回 n! 结果中尾随零的数量。

提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
*/

func trailingZeroes(n int) (res int) {
	for n > 0 {
		n /= 5
		res += n
	}

	return
}

// func trailingZeroes(n int) int {
// 	if n < 5 {
// 		return 0
// 	}
// 	tbl := make(map[int]int)
// 	ans := 0
// 	for i := 5; i <= n; i++ {
// 		if i%10 == 0 {
// 			tbl[i] = tbl[i/10] + 1
// 		} else if i%5 == 0 {
// 			tbl[i] = tbl[i/5] + 1
// 		}
// 		ans += tbl[i]
// 	}
// 	return ans
// }

// func trailingZeroes(n int) int {
// 	if n < 5 {
// 		return 0
// 	}
// 	t, c := n, 0
// 	for t >= 10 && t%10 == 0 {
// 		c++
// 		t /= 10
// 	}
// 	for t >= 5 && t%5 == 0 {
// 		c++
// 		t /= 5
// 	}
// 	return c + trailingZeroes(n-1)
// }
