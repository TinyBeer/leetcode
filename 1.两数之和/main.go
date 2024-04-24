package main

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

// 你可以按任意顺序返回答案。

// 哈希表
func twoSum(nums []int, target int) []int {
	h := make(map[int]int)
	for i, v := range nums {
		j, has := h[target-v]
		if has {
			return []int{j, i}
		}
		h[v] = i
	}
	return nil
}

// 暴力算法
// func twoSum(nums []int, target int) []int {
// 	for i, v := range nums {
// 		for j, m := range nums[i+1:] {
// 			if v+m == target {
// 				return []int{i, i + j + 1}
// 			}
// 		}
// 	}
// 	return nil
// }
