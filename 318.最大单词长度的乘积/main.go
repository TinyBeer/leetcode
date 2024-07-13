package main

/*

给你一个字符串数组 words ，找出并返回 length(words[i]) * length(words[j]) 的最大值，并且这两个单词不含有公共字母。如果不存在这样的两个单词，返回 0 。
*/

func maxProduct(words []string) (ans int) {
	masks := map[int]int{}
	for _, word := range words {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		if len(word) > masks[mask] {
			masks[mask] = len(word)
		}
	}

	for x, lenX := range masks {
		for y, lenY := range masks {
			if x&y == 0 && lenX*lenY > ans {
				ans = lenX * lenY
			}
		}
	}
	return
}

// func maxProduct(words []string) int {
// 	tbl := make([]int, len(words))
// 	for i, w := range words {
// 		for j := 0; j < len(w); j++ {
// 			tbl[i] |= 1 << (w[j] - 'a')
// 		}
// 	}

// 	ans := 0
// 	for i := range words {
// 		for j := i + 1; j < len(words); j++ {
// 			if tbl[i]&tbl[j] == 0 && len(words[i])*len(words[j]) > ans {
// 				ans = len(words[i]) * len(words[j])
// 			}
// 		}
// 	}
// 	return ans
// }

// 暴力解法
// func maxProduct(words []string) int {
// 		ans := 0
// 		for i, w := range words {
// 			tbl := make(map[byte]struct{})
// 			for j := 0; j < len(w); j++ {
// 				tbl[w[j]] = struct{}{}
// 			}
// 		loop:
// 			for j := i + 1; j < len(words); j++ {
// 				for k := 0; k < len(words[j]); k++ {
// 					_, ok := tbl[words[j][k]]
// 					if ok {
// 						continue loop
// 					}
// 				}
// 				ans = max(ans, len(w)*len(words[j]))
// 			}
// 		}
// 		return ans
// }
