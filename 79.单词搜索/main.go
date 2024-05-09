package main

/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。
如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母不允许被重复使用。
*/

// 回溯
func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if m*n < len(word) {
		return false
	}
	tbl := make(map[int]bool)
	var df func(int, int, int) bool
	df = func(i int, j int, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || j < 0 || i >= m || j >= n || tbl[i*n+j] || board[i][j] != word[k] {
			return false
		}
		tbl[i*n+j] = true
		defer func() {
			tbl[i*n+j] = false
		}()
		return df(i, j-1, k+1) || df(i, j+1, k+1) || df(i-1, j, k+1) || df(i+1, j, k+1)
	}
	for i := 0; i < m*n; i++ {
		if df(i/n, i%n, 0) {
			return true
		}
	}
	return false
}
