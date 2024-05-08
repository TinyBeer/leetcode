package main

/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m*n == 0 {
		return m + n
	}
	df := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		df[i] = make([]int, n+1)
		df[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		df[0][j] = j
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				df[i][j] = min(df[i][j-1], df[i-1][j], df[i-1][j-1]-1) + 1
			} else {
				df[i][j] = min(df[i][j-1], df[i-1][j], df[i-1][j-1]) + 1
			}
		}
	}
	return df[m][n]
}
