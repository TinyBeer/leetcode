package main

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

*/
func isValid(s string) bool {
	st := NewStack(len(s))
	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			st.Push(s[i])
		case ')':
			if st.Top() == '(' {
				st.Pop()
			} else {
				return false
			}
		case ']':
			if st.Top() == '[' {
				st.Pop()
			} else {
				return false
			}
		case '}':
			if st.Top() == '{' {
				st.Pop()
			} else {
				return false
			}

		}
	}
	return st.IsEmpty()
}

func NewStack(size int) *Stack {
	return &Stack{
		top: -1,
		arr: make([]byte, size),
	}
}

type Stack struct {
	top int
	arr []byte
}

func (s *Stack) Top() byte {
	if s.top < 0 {
		return '#'
	}

	return s.arr[s.top]
}

func (s *Stack) Push(x byte) {
	s.top++
	s.arr[s.top] = x
}

func (s *Stack) Pop() {
	s.top--
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}
