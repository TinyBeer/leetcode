package main

/*
给你一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。

你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。
*/

func singleNumber(nums []int) []int {
	tmp := 0
	for _, n := range nums {
		tmp ^= n
	}

	lsb := tmp & -tmp
	ans := make([]int, 2)
	for _, n := range nums {
		if n&lsb == 0 {
			ans[0] ^= n
		} else {
			ans[1] ^= n
		}
	}
	return ans
}

// func singleNumber(nums []int) []int {
// 	tbl := make(map[int]int)
// 	for _, v := range nums {
// 		tbl[v]++
// 	}
// 	var ans []int
// 	for k, v := range tbl {
// 		if v == 1 {
// 			ans = append(ans, k)
// 		}
// 	}
// 	return ans
// }
