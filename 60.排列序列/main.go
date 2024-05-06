package main

import (
	"strconv"
)

/*

给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
n = 3	k
1 2 3	1
1 3 2	2
2 1 3	3
2 3 1	4
3 1 2	5
3 2 1	6

*/

// 递归
func getPermutation(n int, k int) string {
	var numstr []string
	for i := 1; i <= n; i++ {
		numstr = append(numstr, strconv.Itoa(i))
	}
	var df func(int, int) string
	k -= 1
	df = func(x int, y int) string {
		if x == n {
			idx := k / y
			k %= y
			s := numstr[idx]
			numstr = append(numstr[:idx], numstr[idx+1:]...)
			return s
		}
		head := df(x+1, y*x)
		idx := k / y
		k %= y
		s := numstr[idx]
		numstr = append(numstr[:idx], numstr[idx+1:]...)
		return head + s
	}
	return df(1, 1)
}
