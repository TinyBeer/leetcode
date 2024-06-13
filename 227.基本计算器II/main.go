package main

/*
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。

整数除法仅保留整数部分。

你可以假设给定的表达式总是有效的。所有中间结果将在 [-231, 231 - 1] 的范围内。

注意：不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
*/

func calculate(s string) int {
	ans := 0
	stack := []int{}
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return ans
}

// func calculate(s string) int {
// 	var ns, os []int
// 	for i := 0; i < len(s); {
// 		switch s[i] {
// 		case ' ':
// 			i++
// 		case '+':
// 			for len(os) > 0 {
// 				l := len(ns)
// 				x, y := ns[l-1], ns[l-2]
// 				ns = ns[:l-2]
// 				op := os[len(os)-1]
// 				os = os[:len(os)-1]
// 				switch op {
// 				case 1:
// 					ns = append(ns, x+y)
// 				case 2:
// 					ns = append(ns, y-x)
// 				case 3:
// 					ns = append(ns, x*y)
// 				case 4:
// 					ns = append(ns, y/x)
// 				}
// 			}
// 			os = append(os, 1)
// 			i++
// 		case '-':
// 			for len(os) > 0 {
// 				l := len(ns)
// 				x, y := ns[l-1], ns[l-2]
// 				ns = ns[:l-2]
// 				op := os[len(os)-1]
// 				os = os[:len(os)-1]
// 				switch op {
// 				case 1:
// 					ns = append(ns, x+y)
// 				case 2:
// 					ns = append(ns, y-x)
// 				case 3:
// 					ns = append(ns, x*y)
// 				case 4:
// 					ns = append(ns, y/x)
// 				}
// 			}
// 			os = append(os, 2)
// 			i++
// 		case '*':
// 			if len(os) != 0 {
// 				l := len(ns)
// 				op := os[len(os)-1]
// 				os = os[:len(os)-1]
// 				switch op {
// 				case 1:
// 					os = append(os, 1)
// 				case 2:
// 					os = append(os, 2)
// 				case 3:
// 					x, y := ns[l-1], ns[l-2]
// 					ns = ns[:l-2]
// 					ns = append(ns, x*y)
// 				case 4:
// 					x, y := ns[l-1], ns[l-2]
// 					ns = ns[:l-2]
// 					ns = append(ns, y/x)
// 				}
// 			}
// 			os = append(os, 3)
// 			i++
// 		case '/':
// 			if len(os) != 0 {
// 				l := len(ns)
// 				op := os[len(os)-1]
// 				os = os[:len(os)-1]
// 				switch op {
// 				case 1:
// 					os = append(os, 1)
// 				case 2:
// 					os = append(os, 2)
// 				case 3:
// 					x, y := ns[l-1], ns[l-2]
// 					ns = ns[:l-2]
// 					ns = append(ns, x*y)
// 				case 4:
// 					x, y := ns[l-1], ns[l-2]
// 					ns = ns[:l-2]
// 					ns = append(ns, y/x)
// 				}
// 			}
// 			os = append(os, 4)
// 			i++
// 		default:
// 			n := 0
// 			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
// 				n = n*10 + int(s[i]-'0')
// 				i++
// 			}
// 			ns = append(ns, n)
// 		}
// 	}

// 	for len(os) > 0 {
// 		l := len(ns)
// 		x, y := ns[l-1], ns[l-2]
// 		ns = ns[:l-2]
// 		op := os[len(os)-1]
// 		os = os[:len(os)-1]
// 		switch op {
// 		case 1:
// 			ns = append(ns, x+y)
// 		case 2:
// 			ns = append(ns, y-x)
// 		case 3:
// 			ns = append(ns, x*y)
// 		case 4:
// 			ns = append(ns, y/x)
// 		}

// 	}
// 	return ns[0]
// }
