package main

/*
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
*/
// kmp算法
func strStr(haystack string, needle string) int {
	// 分别读取原串和匹配串的长度
	n, m := len(haystack), len(needle)
	// 原串和匹配串前面都加空格，使其下标从 1 开始
	s := " " + haystack
	p := " " + needle

	// 构建 next 数组，数组长度为匹配串的长度（next 数组是和匹配串相关的）
	next := make([]int, m+1)
	// 构造过程 i = 2，j = 0 开始，i 小于等于匹配串长度 【构造 i 从 2 开始】
	j := 0
	for i := 2; i <= m; i++ {
		// 匹配不成功的话，j = next(j)
		for j > 0 && p[i] != p[j+1] {
			j = next[j]
		}
		// 匹配成功的话，先让 j++
		if p[i] == p[j+1] {
			j++
		}
		// 更新 next[i]，结束本次循环，i++
		next[i] = j
	}

	// 匹配过程，i = 1，j = 0 开始，i 小于等于原串长度 【匹配 i 从 1 开始】
	j = 0
	for i := 1; i <= n; i++ {
		// 匹配不成功 j = next(j)
		for j > 0 && s[i] != p[j+1] {
			j = next[j]
		}
		// 匹配成功的话，先让 j++，结束本次循环后 i++
		if s[i] == p[j+1] {
			j++
		}
		// 整一段匹配成功，直接返回下标
		if j == m {
			return i - m
		}
	}
	return -1
}

// 暴力解法
// func strStr(haystack string, needle string) int {
// loop:
// 	for i := range haystack {
// 		if len(haystack)-i < len(needle) {
// 			return -1
// 		}
// 		for j := range needle {
// 			if haystack[i+j] != needle[j] {
// 				continue loop
// 			}
// 		}
// 		return i
// 	}
// 	return -1
// }
