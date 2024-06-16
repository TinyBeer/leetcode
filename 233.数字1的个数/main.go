package main

/*
给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
*/

func countDigitOne(n int) (ans int) {
	// mulk 表示 10^k
	// 在下面的代码中，可以发现 k 并没有被直接使用到（都是使用 10^k）
	// 但为了让代码看起来更加直观，这里保留了 k
	for k, mulk := 0, 1; n >= mulk; k++ {
		ans += (n/(mulk*10))*mulk + min(max(n%(mulk*10)-mulk+1, 0), mulk)
		mulk *= 10
	}
	return
}

// func countDigitOne(n int) int {
// 	ans := 0
// 	helper := 1

// 	for n >= helper {
// 		ans += n / (helper * 10) * helper

// 		if n%(helper*10) >= helper {
// 			if n%(helper*10) >= helper*2-1 {
// 				ans += helper
// 			} else {
// 				ans += n%helper + 1
// 			}
// 		}

// 		helper *= 10
// 	}
// 	return ans
// }
