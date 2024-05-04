package main

/*
实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。
*/

func myPow(x float64, n int) float64 {
	ans := 1.0
	flag := n < 0
	po := x
	if flag {
		n = -n
		po = 1 / x
	}
	for ; n != 0; n >>= 1 {
		if n&1 == 1 {
			ans *= po
		}
		po *= po
	}
	return ans
}
