package main

/*
给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。

 s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。

例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"， "cdefab"，"efabcd"， 和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。
返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。

*/
func findSubstring(s string, words []string) []int {
	ls, w, n := len(s), len(words), len(words[0])
	var ans []int
	for i := 0; i < n && i+w*n <= ls; i++ {
		diff := make(map[string]int)
		for j := 0; j < w; j++ {
			diff[s[i+j*n:i+(j+1)*n]]++
		}

		for _, v := range words {
			diff[v]--
			if diff[v] == 0 {
				delete(diff, v)
			}
		}

		for start := i; start < ls-w*n+1; start += n {
			if start != i {
				word := s[start+(w-1)*n : start+w*n]
				diff[word]++
				if diff[word] == 0 {
					delete(diff, word)
				}
				word = s[start-n : start]
				diff[word]--
				if diff[word] == 0 {
					delete(diff, word)
				}
			}
			if len(diff) == 0 {
				ans = append(ans, start)
			}
		}

	}
	return ans
}
