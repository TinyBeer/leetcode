package main

import (
	"strconv"
	"strings"
)

/*

累加数 是一个字符串，组成它的数字可以形成累加序列。

一个有效的 累加序列 必须 至少 包含 3 个数。除了最开始的两个数以外，序列中的每个后续数字必须是它之前两个数字之和。

给你一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是 累加数 。如果是，返回 true ；否则，返回 false 。

说明：累加序列里的数，除数字 0 之外，不会 以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。
*/

func isAdditiveNumber(num string) bool {
	n := len(num)
	var tmp []int
	var check func(string) bool
	check = func(s string) bool {
		if len(s) == 0 {
			return true
		}
		l := len(tmp)
		tmp = append(tmp, tmp[l-1]+tmp[l-2])
		ns := strconv.Itoa(tmp[l])
		return strings.HasPrefix(s, ns) && check(s[len(ns):])
	}
	for i := 1; i <= n/2; i++ {
		if i > 1 && num[0] == '0' {
			break
		}
		n1, _ := strconv.Atoi(num[:i])
		for j := 1; i+j < n; j++ {
			if j > 1 && num[i] == '0' {
				break
			}
			n2, _ := strconv.Atoi(num[i : i+j])
			tmp = append(tmp, n1, n2)
			if check(num[i+j:]) {
				return true
			}
			tmp = nil
		}
	}
	return false
}
