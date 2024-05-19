package main

/*
按字典 wordList 完成从单词 beginWord 到单词 endWord 转化，一个表示此过程的 转换序列 是形式上像 beginWord -> s1 -> s2 -> ... -> sk 这样的单词序列，并满足：

每对相邻的单词之间仅有单个字母不同。
转换过程中的每个单词 si（1 <= i <= k）必须是字典 wordList 中的单词。注意，beginWord 不必是字典 wordList 中的单词。
sk == endWord
给你两个单词 beginWord 和 endWord ，以及一个字典 wordList 。请你找出并返回所有从 beginWord 到 endWord 的 最短转换序列 ，如果不存在这样的转换序列，
返回一个空列表。每个序列都应该以单词列表 [beginWord, s1, s2, ..., sk] 的形式返回。
*/

// 广度优先
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	tbl := make(map[string]bool)
	check := false
	for _, w := range wordList {
		tbl[w] = true
		if w == endWord {
			check = true
		}
	}
	if !check {
		return nil
	}

	res := make([][]string, 0, 1024)
	q := [][]string{{beginWord}}
loop:
	for len(q) > 0 {
		for size := len(q); size > 0; size-- {
			wl := q[0]
			q = q[1:]

			if len(res) != 0 && len(res[0]) <= len(wl) {
				break loop
			}
			// 遍历替换一个字母
			hmap := make(map[string]bool, len(wl))
			for _, v := range wl {
				hmap[v] = true
			}
			w := []byte(wl[len(wl)-1])
			for p := 0; p < len(w); p++ {
				for off := byte(0); off < 26; off++ {
					if w[p] == 'a'+off {
						continue
					}
					nw := append([]byte{}, w...)
					nw[p] = 'a' + off
					if !tbl[string(nw)] || hmap[string(nw)] {
						continue
					}
					nwl := append(wl, string(nw))
					if endWord == string(nw) {
						if len(res) != 0 && len(res[0]) > len(nwl) {
							res = nil
						}
						res = append(res, append([]string{}, nwl...))
						continue
					}
					q = append(q, append([]string{}, nwl...))
				}
			}
		}
	}
	return res
}

// // 暴力探索
// func findLadders(beginWord string, endWord string, wordList []string) [][]string {
// 	check := func() bool {
// 		for _, w := range wordList {
// 			if w == endWord {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	if !check() {
// 		return nil
// 	}

// 	var res [][]string
// 	tmp := make([]string, 0)
// 	tbl := make(map[string]bool)
// 	tbl[beginWord] = true
// 	tmp = append(tmp, beginWord)
// 	var df func(cur string)
// 	df = func(cur string) {
// 		if cur == endWord {
// 			if len(res) != 0 && len(res[0]) < len(tmp) {
// 				return
// 			}
// 			if len(res) != 0 && len(res[0]) > len(tmp) {
// 				res = nil
// 			}
// 			res = append(res, append([]string{}, tmp...))
// 			return
// 		}

// 		for _, w := range wordList {
// 			if tbl[w] || !close(cur, w) {
// 				continue
// 			}
// 			tbl[w] = true
// 			tmp = append(tmp, w)
// 			df(w)
// 			tmp = tmp[:len(tmp)-1]
// 			tbl[w] = false
// 		}

// 	}
// 	df(beginWord)
// 	return res
// }

func isClose(s1, s2 string) bool {
	cnt := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}
		if cnt != 0 {
			return false
		}
		cnt++
	}
	return true
}
