package main

import (
	"strconv"
)

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。
*/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	tbl := make(map[int]int)
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			x := num1[i] - '0'
			y := num2[j] - '0'
			tbl[i+j] += int(x * y)
		}
	}
	sum := 0
	res := ""
	for i := len(num1) + len(num2) - 2; i >= 0; i-- {
		sum /= 10
		sum += int(tbl[i])
		res = strconv.Itoa(sum%10) + res
	}

	sum /= 10
	if sum != 0 {
		res = strconv.Itoa(sum) + res
	}
	return res
}
