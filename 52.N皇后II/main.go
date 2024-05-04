package main

/*
n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。

*/

func totalNQueens(n int) int {
	vertial := make([]bool, n)
	inclined1 := make([]bool, 2*n-1)
	inclined2 := make([]bool, 2*n-1)

	var ans int
	var df func(int)
	df = func(i int) {
		if i == n {
			ans++
			return
		}
		for j := 0; j < n; j++ {
			if vertial[j] || inclined1[n-i+j-1] || inclined2[i+j] {
				continue
			}
			inclined1[n-i+j-1] = true
			vertial[j] = true
			inclined2[i+j] = true
			df(i + 1)
			inclined2[i+j] = false
			vertial[j] = false
			inclined1[n-i+j-1] = false
		}
	}
	df(0)

	return ans
}
