package main

/*
给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回 该列名称对应的列序号 。

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

func titleToNumber(columnTitle string) int {
	ans := 0
	for _, v := range columnTitle {
		ans *= 26
		ans += int(v - 'A' + 1)
	}
	return ans
}
