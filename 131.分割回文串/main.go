package main

/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是
回文串
 。返回 s 所有可能的分割方案。
*/

// 记忆搜索
func partition(s string) [][]string {
	n := len(s)
	f := make([][]int8, n)
	for i := range f {
		f[i] = make([]int8, n)
		f[i][i] = 1
	}
	var isPalindrome func(int, int) int8
	isPalindrome = func(l, r int) int8 {
		if l >= r {
			return 1
		}
		if f[l][r] != 0 {
			return f[l][r]
		}
		f[l][r] = -1
		if s[l] == s[r] {
			f[l][r] = isPalindrome(l+1, r-1)
		}
		return f[l][r]
	}

	var ans [][]string
	var tmp []string
	var df func(i int)
	df = func(l int) {
		if l == n {
			ans = append(ans, append([]string{}, tmp...))
			return
		}
		for r := l; r < n; r++ {
			if isPalindrome(l, r) > 0 {
				tmp = append(tmp, s[l:r+1])
				df(r + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	df(0)
	return ans
}

// // 暴力搜索
// func partition(s string) [][]string {
// 	if len(s) == 1 {
// 		return [][]string{{s}}
// 	}
// 	f := make([][]bool, len(s))
// 	for i := range f {
// 		f[i] = make([]bool, len(s))
// 		for j := range f[i] {
// 			f[i][j] = true
// 		}
// 	}

// 	for i := len(s) - 1; i >= 0; i-- {
// 		for j := i + 1; j < len(s); j++ {
// 			f[i][j] = s[i] == s[j] && f[i+1][j-1]
// 		}
// 	}

// 	var ans [][]string
// 	var tmp []string
// 	var dfs func(int)
// 	dfs = func(i int) {
// 		if i == len(s) {
// 			ans = append(ans, append([]string(nil), tmp...))
// 			return
// 		}
// 		for j := i; j < len(s); j++ {
// 			if f[i][j] {
// 				tmp = append(tmp, s[i:j+1])
// 				dfs(j + 1)
// 				tmp = tmp[:len(tmp)-1]
// 			}
// 		}
// 	}
// 	dfs(0)
// 	return ans
// }
