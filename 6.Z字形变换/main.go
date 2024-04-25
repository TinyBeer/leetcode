package main

import (
	"strings"
)

/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
*/
func convert(s string, numRows int) string {
	if numRows == 1 || numRows > len(s) {
		return s
	}
	tbl := make([]string, numRows)
	for i, r := range s {
		l := i % (numRows*2 - 2)
		if l >= numRows {
			l = numRows*2 - l - 2
		}
		tbl[l] += string(r)
	}
	return strings.Join(tbl, "")
}
