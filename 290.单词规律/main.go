package main

import (
	"strings"
)

/*
给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。
*/

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	ch2w := make(map[byte]string)
	w2ch := make(map[string]byte)
	for i, word := range words {
		ch := pattern[i]
		if w2ch[word] > 0 && w2ch[word] != ch || ch2w[ch] != "" && ch2w[ch] != word {
			return false
		}
		w2ch[word] = ch
		ch2w[ch] = word

	}
	return true
}
