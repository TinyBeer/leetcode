package main

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
*/
func generateParenthesis(n int) []string {
	res := []string{}
	if n == 0 {
		return res
	}
	var recusion func(s string, l int, r int)
	recusion = func(s string, l int, r int) {
		if l == n && r == n {
			res = append(res, s)
			return
		}
		if l < n {
			recusion(s+"(", l+1, r)
		}
		if l > r {
			recusion(s+")", l, r+1)
		}
	}
	recusion("", 0, 0)
	return res
}
