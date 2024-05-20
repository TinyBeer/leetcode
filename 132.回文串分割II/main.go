package main

import (
	"math"
)

/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是
回文串
。

返回符合要求的 最少分割次数 。
*/

func minCut(s string) int {
	n := len(s)
	if n <= 1 {
		return 0
	}
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	g := make([]int, n)
	for i := range f {
		if f[0][i] {
			continue
		}
		g[i] = math.MaxInt
		for j := 0; j < i; j++ {
			if f[j+1][i] && g[j]+1 < g[i] {
				g[i] = g[j] + 1
			}
		}
	}
	return g[n-1]
}
