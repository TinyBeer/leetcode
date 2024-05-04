package main

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
*/

// 计数
func groupAnagrams(strs []string) [][]string {
	tbl := make(map[[26]int][]string)
	for _, s := range strs {
		var cnt [26]int
		for _, c := range s {
			cnt[int(c-'a')]++
		}

		tbl[cnt] = append(tbl[cnt], s)
	}
	var res [][]string
	for _, v := range tbl {
		res = append(res, v)
	}
	return res
}

// // 暴力解法
// func groupAnagrams(strs []string) [][]string {
// 	if len(strs) < 2 {
// 		return append([][]string(nil), strs)
// 	}

// 	tbl := make(map[string][]string)
// 	for _, s := range strs {
// 		bs := []byte(s)
// 		sort.Slice(bs, func(i, j int) bool {
// 			return bs[i] < bs[j]
// 		})
// 		tbl[string(bs)] = append(tbl[string(bs)], s)
// 	}
// 	var res [][]string
// 	for _, v := range tbl {
// 		res = append(res, v)
// 	}
// 	return res
// }
