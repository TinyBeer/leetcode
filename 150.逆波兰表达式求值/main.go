package main

import (
	"strconv"
)

/*
给你一个字符串数组 tokens ，表示一个根据 逆波兰表示法 表示的算术表达式。

请你计算该表达式。返回一个表示表达式值的整数。

注意：

有效的算符为 '+'、'-'、'*' 和 '/' 。
每个操作数（运算对象）都可以是一个整数或者另一个表达式。
两个整数之间的除法总是 向零截断 。
表达式中不含除零运算。
输入是一个根据逆波兰表示法表示的算术表达式。
答案及所有中间计算结果可以用 32 位 整数表示。
*/

type stack struct {
	arr []int
}

func (s *stack) push(x int) {
	s.arr = append(s.arr, x)
}

func (s *stack) pop() int {
	x := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return x
}

func evalRPN(tokens []string) int {
	st := &stack{}
	for _, s := range tokens {
		switch s {
		case "+":
			a, b := st.pop(), st.pop()
			st.push(a + b)
		case "-":
			a, b := st.pop(), st.pop()
			st.push(b - a)
		case "*":
			a, b := st.pop(), st.pop()
			st.push(a * b)
		case "/":
			a, b := st.pop(), st.pop()
			st.push(b / a)
		default:
			n, _ := strconv.Atoi(s)
			st.push(n)
		}

	}
	return st.pop()
}
