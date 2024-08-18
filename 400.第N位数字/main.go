package main

import (
	"math"
	"sort"
)

/*
给你一个整数 n ，请你在无限的整数序列 [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...] 中找出并返回第 n 位上的数字。
*/

func totalDigital(length int) int {
	digits := 0
	for cl, cc := 1, 9; cl <= length; cl++ {
		digits += cl * cc
		cc *= 10
	}
	return digits
}

func findNthDigit(n int) int {
	d := 1 + sort.Search(8, func(length int) bool {
		return totalDigital(length+1) >= n
	})
	preDigits := totalDigital(d - 1)
	idx := n - preDigits - 1
	start := int(math.Pow10(d - 1))
	num := start + idx/d
	digitIdx := idx % d
	return num / int(math.Pow10(d-digitIdx-1)) % 10
}
