package main

/*
DNA序列 由一系列核苷酸组成，缩写为 'A', 'C', 'G' 和 'T'.。

例如，"ACGAATTCCG" 是一个 DNA序列 。
在研究 DNA 时，识别 DNA 中的重复序列非常有用。

给定一个表示 DNA序列 的字符串 s ，返回所有在 DNA
分子中出现不止一次的 长度为 10 的序列(子字符串)。你可以按 任意顺序 返回答案。
*/

const L = 10

var bin = map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

func findRepeatedDnaSequences(s string) (ans []string) {
	n := len(s)
	if n <= L {
		return
	}
	x := 0
	for _, ch := range s[:L-1] {
		x = x<<2 | bin[byte(ch)]
	}
	cnt := map[int]int{}
	for i := 0; i <= n-L; i++ {
		x = (x<<2 | bin[s[i+L-1]]) & (1<<(L*2) - 1)
		cnt[x]++
		if cnt[x] == 2 {
			ans = append(ans, s[i:i+L])
		}
	}
	return ans
}

// func findRepeatedDnaSequences(s string) []string {
// 	if len(s) <= 10 {
// 		return nil
// 	}
// 	tmp := s[:10]
// 	tbl := make(map[string]int)
// 	tbl[tmp]++
// 	for i := 10; i < len(s); i++ {
// 		tmp = tmp[1:] + string(s[i])
// 		tbl[tmp]++
// 	}
// 	var ans []string
// 	for str, count := range tbl {
// 		if count > 1 {
// 			ans = append(ans, str)
// 		}
// 	}
// 	return ans
// }
