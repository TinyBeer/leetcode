package main

import (
	"strconv"
)

/*
给定一个  无重复元素 的 有序 整数数组 nums 。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

"a->b" ，如果 a != b
"a" ，如果 a == b
*/

func summaryRanges(nums []int) (ans []string) {
	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return
}

// func summaryRanges(nums []int) []string {
// 	var res []string
// 	last := 0
// 	i := 0
// 	for i < len(nums) {
// 		for i+1 < len(nums) && nums[i]+1 == nums[i+1] {
// 			i++
// 		}
// 		if last == i {
// 			res = append(res, strconv.Itoa(nums[i]))
// 		} else {
// 			res = append(res, strconv.Itoa(nums[last])+"->"+strconv.Itoa(nums[i]))
// 		}
// 		i++
// 		last = i
// 	}
// 	return res
// }
