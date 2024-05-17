package main

/*
给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。
*/

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}
	res := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		res[i] = 1
		for j := i - 1; j >= 1; j-- {
			res[j] = res[j] + res[j-1]
		}
	}
	return res
}
