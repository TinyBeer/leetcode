package main

/*
给定一个正整数 n ，你可以做如下操作：

如果 n 是偶数，则用 n / 2替换 n 。
如果 n 是奇数，则可以用 n + 1或n - 1替换 n 。
返回 n 变为 1 所需的 最小替换次数 。
*/

// 111101  111101
// 111100  111110
// 11110   11111
// 1111    100000
// 10000   10000
// 1000    1000

func integerReplacement(n int) int {
	ans := 0
	for n != 1 {
		if n&1 == 1 {
			if n == 3 || n&0x03 != 0x03 {
				n -= 1
			} else {
				n += 1
			}
		} else {
			n >>= 1
		}
		ans++
	}
	return ans
}
