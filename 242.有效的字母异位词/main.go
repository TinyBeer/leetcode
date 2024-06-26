package main

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	tbl := make(map[rune]int)
	for _, r := range s {
		tbl[r]++
	}
	for _, r := range t {
		if tbl[r] == 0 {
			return false
		}
		tbl[r]--
	}
	return true
}
