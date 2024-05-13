package main

/*
给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。

两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空
子字符串
：

s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
注意：a + b 意味着字符串 a 和 b 连接。
*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	if m*n == 0 {
		return s1+s2 == s3
	}
	df := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		df[i] = make([]bool, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 {
				if j == 0 {
					df[i][j] = true
				} else {
					df[i][j] = df[i][j-1] && s2[j-1] == s3[j-1]
				}
			} else if j == 0 {
				df[i][j] = df[i-1][j] && s1[i-1] == s3[i-1]
			} else {
				df[i][j] = (df[i][j-1] && s2[j-1] == s3[i+j-1]) || (df[i-1][j] && s1[i-1] == s3[i+j-1])
			}
		}
	}
	return df[m][n]
}

// // 暴力接
// func isInterleave(s1 string, s2 string, s3 string) bool {
// 	if len(s1)+len(s2) != len(s3) {
// 		return false
// 	}
// 	var df func(i, j, k int) bool
// 	df = func(i, j, k int) bool {

// 		if k == len(s3) {
// 			return true
// 		}
// 		if i < len(s1) && s1[i] == s3[k] && df(i+1, j, k+1) {
// 			return true
// 		}

// 		if j < len(s2) && s2[j] == s3[k] && df(i, j+1, k+1) {
// 			return true
// 		}
// 		return false
// 	}
// 	return df(0, 0, 0)
// }
