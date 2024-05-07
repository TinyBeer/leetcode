package main

/*
给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
*/

func addBinary(a string, b string) string {
	tbl := map[int]string{0: "0", 1: "1", 2: "0", 3: "1"}
	carry := 0
	ans := ""
	for i := 1; i <= len(a) || i <= len(b); i++ {
		if i <= len(a) {
			carry += int(a[len(a)-i] - '0')
		}
		if i <= len(b) {
			carry += int(b[len(b)-i] - '0')
		}
		ans = tbl[carry] + ans
		carry /= 2
	}
	if carry > 0 {
		ans = tbl[carry] + ans
	}
	return ans
}
