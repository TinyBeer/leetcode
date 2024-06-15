package main

/*
给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。

如果存在一个整数 x 使得 n == 2x ，则认为 n 是 2 的幂次方。
*/

func isPowerOfTwo(n int) bool {
	n = n&0x55555555 + n>>1&0x55555555
	n = n&0x33333333 + n>>2&0x33333333
	n = n&0x0f0f0f0f + n>>4&0x0f0f0f0f
	n = n&0x00ff00ff + n>>8&0x00ff00ff
	n = n&0x0000ffff + n>>16&0x0000ffff
	return n == 1
}
