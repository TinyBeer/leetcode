package main

/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
*/

// stake
func longestValidParentheses(s string) int {
	maxAns := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// dp
// func longestValidParentheses(s string) int {
// 	ans := 0
// 	dp := make([]int, len(s))
// 	for i := 1; i < len(s); i++ {
// 		if s[i] == ')' {
// 			if s[i-1] == '(' {
// 				if i >= 2 {
// 					dp[i] = dp[i-2] + 2
// 				} else {
// 					dp[i] = 2
// 				}
// 			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
// 				if i-dp[i-1] >= 2 {
// 					dp[i] = dp[i-dp[i-1]-2] + dp[i-1] + 2
// 				} else {
// 					dp[i] = dp[i-1] + 2
// 				}
// 			}
// 		}
// 		if ans < dp[i] {
// 			ans = dp[i]
// 		}
// 	}
// 	return ans
// }

// 暴力解法
// func longestValidParentheses(s string) int {
// 	ans := 0
// 	for i := range s {
// 		cnt := 0
// 		for cur := i; cur < len(s); cur++ {
// 			if s[cur] == '(' {
// 				cnt++
// 			} else {
// 				cnt--
// 			}

// 			if cnt < 0 {
// 				break
// 			}
// 			if cnt == 0 && ans < cur-i+1 {
// 				ans = cur - i + 1
// 			}
// 		}
// 	}
// 	return ans
// }
