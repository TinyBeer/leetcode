package main

/*
给定一个仅包含数字 0-9 的字符串 num 和一个目标值整数 target ，在 num 的数字之间添加 二元 运算符（不是一元）+、- 或 * ，返回 所有 能够得到 target 的表达式。

注意，返回表达式中的操作数 不应该 包含前导零。
*/

func addOperators(num string, target int) []string {
	n := len(num)
	var ans []string
	var backtrack func(expr []byte, cur, res, mul int)
	backtrack = func(expr []byte, cur, res, mul int) {
		if cur == n {
			if res == target {
				ans = append(ans, string(expr))
			}
			return
		}
		signIdx := len(expr)
		if cur > 0 {
			expr = append(expr, 0)
		}
		for i, val := cur, 0; i < n && (cur == i || num[cur] != '0'); i++ {
			val = val*10 + int(num[i]-'0')
			expr = append(expr, num[i])
			if cur == 0 {
				backtrack(expr, i+1, val, val)
			} else {
				expr[signIdx] = '+'
				backtrack(expr, i+1, res+val, val)
				expr[signIdx] = '-'
				backtrack(expr, i+1, res-val, -val)
				expr[signIdx] = '*'
				backtrack(expr, i+1, res-mul+mul*val, mul*val)
			}
		}
	}
	backtrack(make([]byte, 0, 2*n-1), 0, 0, 0)
	return ans
}
