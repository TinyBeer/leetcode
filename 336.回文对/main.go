package main

import "sort"

/*

给定一个由唯一字符串构成的 0 索引 数组 words 。

回文对 是一对整数 (i, j) ，满足以下条件：

0 <= i, j < words.length，
i != j ，并且
words[i] + words[j]（两个字符串的连接）是一个
回文串
。
返回一个数组，它包含 words 中所有满足 回文对 条件的字符串。

你必须设计一个时间复杂度为 O(sum of words[i].length) 的算法。
*/

func palindromePairs(words []string) [][]int {
	ans := [][]int{}
	index := make([]int, len(words))
	for i := range index {
		index[i] = i
	}
	sort.Slice(index, func(i, j int) bool {
		return len(words[index[i]]) < len(words[index[j]])
	})
	// 存储反序字符串
	m := make(map[string]int)
	// kmp
	getNext := func(t string) []int {
		next := make([]int, len(t))
		for i := range next {
			next[i] = -1
		}
		for j := 1; j < len(t); j++ {
			k := next[j-1]
			for k != -1 && t[k+1] != t[j] {
				k = next[k]
			}
			if t[k+1] == t[j] {
				next[j] = k + 1
			}
		}
		return next
	}
	// 从短到长遍历
	for _, i := range index {
		// 正序和反序
		word := words[i]
		rword := reverse(word)
		next := getNext(word)
		rnext := getNext(rword)
		// 前缀回文，word模式串，rword查询串
		best := -1
		for i := 0; i < len(rword); i++ {
			for best != -1 && word[best+1] != rword[i] {
				best = next[best]
			}
			if word[best+1] == rword[i] {
				best += 1
			}
		}
		// 最终得到的best是最长回文前缀的右端索引
		// 而由best = next[best]可以依次递归获得其余的回文前缀的右端索引
		for best != -1 {
			if v, ok := m[word[best+1:]]; ok {
				// 除去回文前缀的部分，匹配到v字符串
				ans = append(ans, []int{v, i})
			}
			best = next[best]
		}

		// 后缀回文，rword模式串，word查询串
		best = -1
		for i := 0; i < len(word); i++ {
			for best != -1 && rword[best+1] != word[i] {
				best = rnext[best]
			}
			if rword[best+1] == word[i] {
				best += 1
			}
		}
		for best != -1 {
			if v, ok := m[word[:len(word)-best-1]]; ok {
				ans = append(ans, []int{i, v})
			}
			best = rnext[best]
		}
		// 考虑回文前缀和后缀为空的情况
		if v, ok := m[word]; ok {
			ans = append(ans, []int{v, i})
			ans = append(ans, []int{i, v})
		}
		// 存储倒序
		m[rword] = i
	}
	return ans
}

func reverse(s string) (t string) {
	n := len(s)
	tmp := make([]byte, n)
	for i := range s {
		tmp[i] = s[n-1-i]
	}
	return string(tmp)
}

// func palindromePairs(words []string) [][]int {
// 	tbl := make(map[string][]int)
// 	for i, w := range words {
// 		tbl[w] = append(tbl[w], i)
// 	}
// 	var res [][]int
// 	for i, w := range words {
// 		for j := 0; j <= len(w); j++ {
// 			if isPalindrome(w[j:]) {
// 				bs := []byte(w[:j])
// 				slices.Reverse(bs)
// 				if idxs, ok := tbl[string(bs)]; ok {
// 					for _, idx := range idxs {
// 						if idx != i {
// 							res = append(res, []int{i, idx})
// 						}
// 					}
// 				}
// 			}
// 			if j > 0 && isPalindrome(w[:j]) {
// 				bs := []byte(w[j:])
// 				slices.Reverse(bs)
// 				if idxs, ok := tbl[string(bs)]; ok {
// 					for _, idx := range idxs {
// 						if idx != i {
// 							res = append(res, []int{idx, i})
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// func isPalindrome(str string) bool {
// 	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
// 		if str[i] != str[j] {
// 			return false
// 		}
// 	}
// 	return true
// }
