package main

/*
给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。

例如：

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...
*/

func convertToTitle(columnNumber int) string {
	var res []byte
	for columnNumber > 0 {
		columnNumber--
		res = append(res, 'A'+byte(columnNumber%26))
		columnNumber /= 26
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
