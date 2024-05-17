package main

/*
给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数，结果需要对 109 + 7 取模。
*/

/*
  i	rabbbit   rabbit
  0 1 1 1 1 1 1 1 1
  1   1 1 1 1 1 1 1
  2     1 1 1 1 1 1
  3       1 2 3 3 3
  4         1 3 3 3
  5           0 3 3
  6             0 3

*/

func numDistinct(s string, t string) int {
	len_s := len(s)
	len_t := len(t)
	if len_s < len_t {
		return 0
	}
	dp := make([][]int, 1+len_s)
	for i := 0; i < len_s+1; i++ {
		dp[i] = make([]int, 1+len_t)
	}
	for i := 0; i < len_s+1; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < len_s+1; i++ {
		for j := 1; j < len_t+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len_s][len_t] % (10e9 + 7)
}
