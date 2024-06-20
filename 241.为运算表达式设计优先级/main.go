package main

import "unicode"

/*
给你一个由数字和运算符组成的字符串 expression ，按不同优先级组合数字和运算符，计算并返回所有可能组合的结果。你可以 按任意顺序 返回答案。

生成的测试用例满足其对应输出值符合 32 位整数范围，不同结果的数量不超过 104 。
*/

const addition, subtraction, multiplication = -1, -2, -3

func diffWaysToCompute(expression string) []int {
	ops := []int{}
	for i, n := 0, len(expression); i < n; {
		if unicode.IsDigit(rune(expression[i])) {
			x := 0
			for ; i < n && unicode.IsDigit(rune(expression[i])); i++ {
				x = x*10 + int(expression[i]-'0')
			}
			ops = append(ops, x)
		} else {
			if expression[i] == '+' {
				ops = append(ops, addition)
			} else if expression[i] == '-' {
				ops = append(ops, subtraction)
			} else {
				ops = append(ops, multiplication)
			}
			i++
		}
	}

	n := len(ops)
	dp := make([][][]int, n)
	for i, x := range ops {
		dp[i] = make([][]int, n)
		dp[i][i] = []int{x}
	}
	for sz := 3; sz <= n; sz++ {
		for l, r := 0, sz-1; r < n; l += 2 {
			for k := l + 1; k < r; k += 2 {
				for _, x := range dp[l][k-1] {
					for _, y := range dp[k+1][r] {
						if ops[k] == addition {
							dp[l][r] = append(dp[l][r], x+y)
						} else if ops[k] == subtraction {
							dp[l][r] = append(dp[l][r], x-y)
						} else {
							dp[l][r] = append(dp[l][r], x*y)
						}
					}
				}
			}
			r += 2
		}
	}
	return dp[0][n-1]
}

// func diffWaysToCompute(expression string) []int {
// 	var ops []int
// 	for i := 0; i < len(expression); {
// 		switch expression[i] {
// 		case '+':
// 			ops = append(ops, -1)
// 			i++
// 		case '-':
// 			ops = append(ops, -2)
// 			i++
// 		case '*':
// 			ops = append(ops, -3)
// 			i++
// 		default:
// 			n := 0
// 			for i < len(expression) && expression[i] >= '0' && expression[i] <= '9' {
// 				n *= 10
// 				n += int(expression[i] - '0')
// 				i++
// 			}
// 			ops = append(ops, n)
// 		}
// 	}

// 	dp := make([][][]int, len(ops))
// 	for i := range dp {
// 		dp[i] = make([][]int, len(ops))
// 	}
// 	var dfs func(l, r int) []int
// 	dfs = func(l, r int) []int {
// 		res := dp[l][r]
// 		if res != nil {
// 			return res
// 		}
// 		if l == r {
// 			dp[l][r] = []int{ops[l]}
// 			return dp[l][r]
// 		}
// 		for i := l; i < r; i += 2 {
// 			left := dfs(l, i)
// 			right := dfs(i+2, r)
// 			for _, x := range left {
// 				for _, y := range right {
// 					switch ops[i+1] {
// 					case -1:
// 						dp[l][r] = append(dp[l][r], x+y)
// 					case -2:
// 						dp[l][r] = append(dp[l][r], x-y)
// 					case -3:
// 						dp[l][r] = append(dp[l][r], x*y)
// 					}
// 				}
// 			}

// 		}
// 		return dp[l][r]
// 	}

// 	return dfs(0, len(ops)-1)
// }
