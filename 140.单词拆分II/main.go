package main

import (
	"strings"
)

/*

给定一个字符串 s 和一个字符串字典 wordDict ，在字符串 s
中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序 返回所有这些可能的句子。

注意：词典中的同一个单词可能在分段中被重复使用多次。
*/

func wordBreak(s string, wordDict []string) []string {

	var res []string
	var tmp []string
	var df func(i int)
	df = func(i int) {
		if i == len(s) {
			res = append(res, strings.Join(tmp, " "))
			return
		}
		for _, w := range wordDict {
			if strings.HasPrefix(s[i:], w) {
				tmp = append(tmp, w)
				df(i + len(w))
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	df(0)
	return res
}
