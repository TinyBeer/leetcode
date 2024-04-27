package main

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
*/

// 逐次探索
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 暴力
// func longestCommonPrefix(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}
// 	if len(strs) == 1 {
// 		return strs[0]
// 	}

// 	res := ""
// loop:
// 	for i := range strs[0] {
// 		for j := 1; j < len(strs); j++ {
// 			if len(strs[j]) <= i || strs[0][i] != strs[j][i] {
// 				break loop
// 			}
// 		}
// 		res += string(strs[0][i])
// 	}
// 	return res
// }
