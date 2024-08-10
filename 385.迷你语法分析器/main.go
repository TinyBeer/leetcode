package main

/*
给定一个字符串 s 表示一个整数嵌套列表，实现一个解析它的语法分析器并返回解析的结果 NestedInteger 。

列表中的每个元素只可能是整数或整数嵌套列表
*/

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	var idx int
	var helper func() NestedInteger
	helper = func() NestedInteger {
		res := NestedInteger{}
		isNegative := false
		isNested := false
		num := 0
	loop:
		for idx < len(s) {
			switch s[idx] {
			case '-':
				isNegative = true
				idx++
				continue
			case '[':
				isNested = true
				var ni NestedInteger
				if s[idx+1] != ']' {
					idx++
					ni = helper()
				} else {
					idx += 2
					return ni

				}
				res.Add(ni)
			case ',':
				if isNested {
					idx++
					ni := helper()
					res.Add(ni)
				} else {
					break loop
				}
			case ']':
				if isNested {
					idx++
					return res
				} else {
					break loop
				}
			default:
				num *= 10
				num += int(s[idx] - '0')
				idx++
			}
		}
		if isNegative {
			num *= -1
		}
		res.SetInteger(num)
		return res
	}
	ni := helper()
	return &ni
}
