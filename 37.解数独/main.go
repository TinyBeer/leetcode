package main

/*
编写一个程序，通过填充空格来解决数独问题。

数独的解法需 遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用 '.' 表示。
*/

func solveSudoku(board [][]byte) {
	var row, col [9][9]byte
	var sub [3][3][9]byte

	for i, r := range board {
		for j, v := range r {
			if v == '.' {
				continue
			}
			n := v - '1'
			row[i][n]++
			col[j][n]++
			sub[i/3][j/3][n]++
		}
	}

	var explore func(int, int) bool
	explore = func(i int, j int) bool {
		if i < 0 {
			return true
		}
		if board[i][j] != '.' {
			return explore(next(i, j))
		}
		for n := 0; n < 9; n++ {
			if row[i][n] == 1 || col[j][n] == 1 || sub[i/3][j/3][n] == 1 {
				continue
			}
			board[i][j] = '1' + byte(n)
			row[i][n]++
			col[j][n]++
			sub[i/3][j/3][n]++
			if explore(next(i, j)) {
				return true
			}
			board[i][j] = '.'
			row[i][n]--
			col[j][n]--
			sub[i/3][j/3][n]--
		}
		return false
	}
	explore(0, 0)
}

func next(i int, j int) (int, int) {
	if j == 8 {
		if i == 8 {
			return -1, -1
		} else {
			j = 0
			i++
		}
	} else {
		j++
	}
	return i, j
}
