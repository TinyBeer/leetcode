package main

import (
	"strings"
)

/*

按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。


1 2 3 4
1 2 3 4
1 2 3 4
1 2 3 4`

*/

func solveNQueens(n int) [][]string {
	vertial := make([]bool, n)
	inclined1 := make([]bool, 2*n-1)
	inclined2 := make([]bool, 2*n-1)

	var res [][]string
	tmp := make([]string, n)
	var df func(int)
	df = func(i int) {
		if i == n {
			res = append(res, append([]string(nil), tmp...))
			return
		}
		for j := 0; j < n; j++ {
			if vertial[j] || inclined1[n-i+j-1] || inclined2[i+j] {
				continue
			}
			inclined1[n-i+j-1] = true
			vertial[j] = true
			inclined2[i+j] = true
			tmp[i] = strings.Repeat(".", j) + "Q" + strings.Repeat(".", n-j-1)
			df(i + 1)
			inclined2[i+j] = false
			vertial[j] = false
			inclined1[n-i+j-1] = false
		}
	}
	df(0)

	return res
}
