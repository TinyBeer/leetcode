package main

/*
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
*/

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	df := make([]bool, len(s)+1)
	df[0] = true
	for i := range df {
		for _, w := range wordDict {
			if i >= len(w) {
				df[i] = df[i] ||
					(s[i-len(w):i] == w &&
						df[i-len(w)])
			}
		}
	}
	return df[len(s)]
}
