package main

/*
给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。
*/

// func firstUniqChar(s string) int {
// 	tbl := make(map[byte]int)
// 	for i := range s {
// 		tbl[s[i]]++
// 	}
// 	ans := -1
// 	for i := range s {
// 		if tbl[s[i]] == 1 {
// 			ans = i
// 			break
// 		}
// 	}
// 	return ans
// }

func firstUniqChar(s string) int {
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		pos[i] = n
	}
	for i, ch := range s {
		ch -= 'a'
		if pos[ch] == n {
			pos[ch] = i
		} else {
			pos[ch] = n + 1
		}
	}
	ans := n
	for _, p := range pos[:] {
		if p < ans {
			ans = p
		}
	}
	if ans < n {
		return ans
	}
	return -1
}
