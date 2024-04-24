package main

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。

func lengthOfLongestSubstring(s string) int {
	h := make(map[rune]int)
	max := 0
	l := 0
	for i, r := range s {
		j, ok := h[r]
		h[r] = i
		l++
		if ok && l > i-j {
			l = i - j
		}

		if max < l {
			max = l
		}
	}
	return max
}
