package main

/*
请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）


注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
空白格用 '.' 表示。
*/

// 暴力
func isValidSudoku(board [][]byte) bool {
	var rows, cols [9][9]int
	var sub [3][3][9]int
	for i, row := range board {
		for j, num := range row {
			if num != '.' {
				index := num - '1'
				rows[i][index]++
				cols[j][index]++
				sub[i/3][j/3][index]++
				if rows[i][index] > 1 || cols[j][index] > 1 || sub[i/3][j/3][index] > 1 {
					return false
				}
			}
		}
	}
	return true
}
