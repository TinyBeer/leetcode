package main

import "strings"

/*
给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。

如果不存在这样的子字符串，则返回 0。
*/

func longestSubstring(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	tbl := [26]int{}
	for _, ch := range s {
		tbl[ch-'a']++
	}
	var split byte
	for i, cnt := range tbl {
		if cnt > 0 && cnt < k {
			split = 'a' + byte(i)
			break
		}
	}
	if split == 0 {
		return len(s)
	}
	var ans int
	for _, subStr := range strings.Split(s, string(split)) {
		ans = max(ans, longestSubstring(subStr, k))
	}
	return ans
}
