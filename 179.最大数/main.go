package main

import (
	"sort"
	"strconv"
)

/*
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
*/

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return x*sy+y > y*sx+x
	})
	if nums[0] == 0 {
		return "0"
	}
	var ans []byte
	for _, v := range nums {
		ans = append(ans, strconv.Itoa(v)...)
	}
	return string(ans)
}
