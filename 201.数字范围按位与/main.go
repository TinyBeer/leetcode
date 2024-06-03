package main

/*
给你两个整数 left 和 right ，表示区间 [left, right] ，
返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）。
*/

func rangeBitwiseAnd(left int, right int) int {
	for i := 0; i < 64; i++ {
		if left == right {
			return left << i
		}
		left >>= 1
		right >>= 1
	}
	return 0
}
