package main

/*
给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。

返回所有可能的结果。答案可以按 任意顺序 返回
*/

func removeInvalidParentheses(s string) []string {
	left, right := 0, 0
	for i := range s {
		switch s[i] {
		case '(':
			left++
		case ')':
			if left == 0 {
				right++
			} else {
				left--
			}
		}
	}

	isValid := func(str string) bool {
		cnt := 0
		for _, ch := range str {
			if ch == '(' {
				cnt++
			} else if ch == ')' {
				cnt--
				if cnt < 0 {
					return false
				}
			}
		}
		return cnt == 0
	}

	var dfs func(ans *[]string, str string, start int, left int, right int)
	dfs = func(ans *[]string, str string, start, left, right int) {
		if left == 0 && right == 0 {
			if isValid(str) {
				*ans = append(*ans, str)
			}
			return
		}
		for i := start; i < len(str); i++ {
			if i != start && str[i] == str[i-1] {
				continue
			}
			if left+right > len(str)-i {
				return
			}

			if left > 0 && str[i] == '(' {
				dfs(ans, str[:i]+str[i+1:], i, left-1, right)
			}

			if right > 0 && str[i] == ')' {
				dfs(ans, str[:i]+str[i+1:], i, left, right-1)
			}
		}
	}
	var ans []string
	dfs(&ans, s, 0, left, right)
	return ans
}
