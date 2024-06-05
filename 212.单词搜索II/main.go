package main

/*
给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。

单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母在一个单词中不允许被重复使用。
*/

type Trie struct {
	children map[byte]*Trie
	word     string
}

func (t *Trie) Insert(word string) {
	node := t
	for i := range word {
		ch := word[i]
		if node.children[ch] == nil {
			node.children[ch] = &Trie{children: map[byte]*Trie{}}
		}
		node = node.children[ch]
	}
	node.word = word
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findWords(board [][]byte, words []string) (ans []string) {
	t := &Trie{children: map[byte]*Trie{}}
	for _, word := range words {
		t.Insert(word)
	}

	m, n := len(board), len(board[0])

	var dfs func(node *Trie, x, y int)
	dfs = func(node *Trie, x, y int) {
		ch := board[x][y]
		nxt := node.children[ch]
		if nxt == nil {
			return
		}

		if nxt.word != "" {
			ans = append(ans, nxt.word)
			nxt.word = ""
		}

		if len(nxt.children) > 0 {
			board[x][y] = '#'
			for _, d := range dirs {
				nx, ny := x+d.x, y+d.y
				if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
					dfs(nxt, nx, ny)
				}
			}
			board[x][y] = ch
		}

		if len(nxt.children) == 0 {
			delete(node.children, ch)
		}
	}
	for i, row := range board {
		for j := range row {
			dfs(t, i, j)
		}
	}

	return
}

// type Trie struct {
// 	children [26]*Trie
// 	word     string
// }

// func (t *Trie) insert(word string) {
// 	node := t
// 	for _, ch := range word {
// 		ch -= 'a'
// 		if node.children[ch] == nil {
// 			node.children[ch] = &Trie{}
// 		}
// 		node = node.children[ch]
// 	}
// 	node.word = word
// }

// func findWords(board [][]byte, words []string) []string {
// 	dirs := []struct{ x, y int }{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
// 	t := &Trie{}
// 	for _, w := range words {
// 		t.insert(w)
// 	}
// 	m, n := len(board), len(board[0])
// 	seen := map[string]bool{}
// 	var dfs func(node *Trie, x, y int)
// 	dfs = func(node *Trie, x, y int) {
// 		ch := board[x][y]
// 		node = node.children[ch-'a']
// 		if node == nil {
// 			return
// 		}

// 		if node.word != "" {
// 			seen[node.word] = true
// 		}

// 		board[x][y] = '#'
// 		for _, d := range dirs {
// 			nx, ny := x+d.x, y+d.y
// 			if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
// 				dfs(node, nx, ny)
// 			}
// 		}
// 		board[x][y] = ch
// 	}
// 	for i, row := range board {
// 		for j := range row {
// 			dfs(t, i, j)
// 		}
// 	}

// 	ans := make([]string, 0, len(seen))
// 	for s := range seen {
// 		ans = append(ans, s)
// 	}
// 	return ans
// }

// 暴力
// func findWords(board [][]byte, words []string) []string {
// 	m, n := len(board), len(board[0])
// 	visted := make([]bool, m*n)
// 	var def func(word string, i, j int) bool
// 	def = func(word string, i, j int) bool {
// 		if i < 0 || j < 0 || i == m || j == n || visted[i*n+j] || board[i][j] != word[0] {
// 			return false
// 		}
// 		if len(word) == 1 {
// 			return true
// 		}
// 		visted[i*n+j] = true
// 		for _, p := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
// 			if def(word[1:], i+p[0], j+p[1]) {
// 				return true
// 			}
// 		}
// 		visted[i*n+j] = false
// 		return false
// 	}
// 	var res []string
// loop:
// 	for _, word := range words {
// 		for i := range board {
// 			for j := range board[i] {
// 				if def(word, i, j) {
// 					res = append(res, word)
// 					visted = make([]bool, m*n)
// 					continue loop
// 				}
// 			}
// 		}
// 	}
// 	return res
// }
