package main

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
*/

/*
正则组合
xy  匹配x
x*  匹配0~n个x
.*  匹配任意字符串

*/

// 动态规划
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}

// 递归
// func isMatch(s string, p string) bool {
// 	if p == ".*" {
// 		return true
// 	}
// 	if len(p) == 0 {
// 		return len(s) == 0
// 	}

// 	if len(s) == 0 {
// 		if len(p)%2 != 0 {
// 			return false
// 		}
// 		for i := 1; i < len(p); i += 2 {
// 			if p[i] != '*' {
// 				return false
// 			}
// 		}
// 		return true
// 	}

// 	if len(p) == 1 {
// 		if p[0] == '.' {
// 			return len(s) == 1
// 		}
// 		return s == p
// 	}

// 	res := false
// 	if p[1] == '*' {
// 		res = isMatch(s, p[2:])
// 		if p[0] == '.' {
// 			for i := range s {
// 				res = res || isMatch(s[i+1:], p[2:])
// 			}
// 		} else {
// 			for i := range s {
// 				if p[0] != s[i] {
// 					break
// 				}
// 				res = res || isMatch(s[i+1:], p[2:])
// 			}
// 		}
// 		return res
// 	}
// 	if p[0] == s[0] || p[0] == '.' {
// 		return isMatch(s[1:], p[1:])
// 	}
// 	return false
// }
