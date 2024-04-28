package main

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	tbl := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	res := []string{}
	var recursion func(pre string, i int)
	recursion = func(pre string, i int) {
		if i == len(digits) {
			res = append(res, pre)
			return
		}
		for _, v := range tbl[digits[i]] {
			recursion(pre+string(v), i+1)
		}
	}
	recursion("", 0)
	return res
}
