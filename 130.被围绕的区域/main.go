package main

/*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，
找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
*/

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	m, n := len(board), len(board[0])
	dir := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	search := func(r int) {
		board[r/n][r%n] = 'A'
		q := []int{r}
		for len(q) > 0 {
			tmp := q[0]
			q = q[1:]
			i, j := tmp/n, tmp%n
			for _, d := range dir {
				nx, ny := i+d[0], j+d[1]
				if nx < 0 || ny < 0 || nx == m || ny == n {
					continue
				}
				if board[nx][ny] == 'O' {
					board[nx][ny] = 'A'
					q = append(q, nx*n+ny)
				}
			}
		}
	}
	for i, j := 0, 0; j < n; j++ {
		if board[i][j] == 'O' {
			search(i*n + j)
		}

	}
	for i, j := 0, 0; i < m; i++ {
		if board[i][j] == 'O' {
			search(i*n + j)
		}

	}
	for i, j := m-1, 0; j < n; j++ {
		if board[i][j] == 'O' {
			search(i*n + j)
		}

	}
	for i, j := 0, n-1; i < m; i++ {
		if board[i][j] == 'O' {
			search(i*n + j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

// func solve(board [][]byte) {
// 	if len(board) == 0 {
// 		return
// 	}

// 	m, n := len(board), len(board[0])
// 	travel, mask := make([]bool, m*n), make([]bool, m*n)

// 	type port struct {
// 		x, y int
// 	}
// 	dir := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
// 	search := func(i, j int) {
// 		q := []port{{i, j}}
// 		if mask[i*n+j] {
// 			return
// 		}
// 		mask[i*n+j] = true
// 		for x := 0; x < len(q); x++ {
// 			p := q[x]
// 			travel[p.x*n+p.y] = true
// 			mask[p.x*n+p.y] = true
// 			for _, d := range dir {
// 				nx, ny := p.x+d[0], p.y+d[1]
// 				if nx < 0 || ny < 0 || nx == m || ny == n {
// 					continue
// 				}
// 				if mask[nx*n+ny] {
// 					continue
// 				}
// 				if board[nx][ny] == 'O' {
// 					q = append(q, port{nx, ny})
// 				}
// 			}
// 		}
// 	}
// 	for i, j := 0, 0; j < n; j++ {
// 		if board[i][j] == 'X' {
// 			continue
// 		}
// 		search(i, j)
// 	}
// 	for i, j := 0, 0; i < m; i++ {
// 		if board[i][j] == 'X' {
// 			continue
// 		}
// 		search(i, j)
// 	}
// 	for i, j := m-1, 0; j < n; j++ {
// 		if board[i][j] == 'X' {
// 			continue
// 		}
// 		search(i, j)
// 	}
// 	for i, j := 0, n-1; i < m; i++ {
// 		if board[i][j] == 'X' {
// 			continue
// 		}
// 		search(i, j)
// 	}
// 	for i, v := range travel {
// 		if !v {
// 			x, y := i/n, i%n
// 			board[x][y] = 'X'
// 		}
// 	}

// }

// func solve(board [][]byte) {
// 	if len(board) == 0 {
// 		return
// 	}

// 	type port struct {
// 		x, y int
// 	}
// 	dir := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
// 	m, n := len(board), len(board[0])
// 	mask := make([]bool, m*n)
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if mask[i*n+j] {
// 				continue
// 			}
// 			mask[i*n+j] = true
// 			if board[i][j] == 'X' {
// 				continue
// 			}

// 			flag := false
// 			q := []port{{i, j}}
// 			for x := 0; x < len(q); x++ {
// 				p := q[x]
// 				mask[p.x*n+p.y] = true
// 				for _, d := range dir {
// 					nx, ny := p.x+d[0], p.y+d[1]
// 					if nx < 0 || ny < 0 || nx == m || ny == n {
// 						flag = true
// 						continue
// 					}
// 					if mask[nx*n+ny] {
// 						continue
// 					}
// 					if board[nx][ny] == 'O' {
// 						q = append(q, port{nx, ny})
// 					}
// 				}
// 			}
// 			if !flag {
// 				for _, p := range q {
// 					board[p.x][p.y] = 'X'
// 				}
// 			}
// 		}
// 	}
// }
