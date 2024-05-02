package main

/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
*/

// 置换法
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

// // 哈希表
// func firstMissingPositive(nums []int) int {
// 	tbl := make(map[int]struct{})
// 	min := 1
// 	for _, n := range nums {
// 		tbl[n] = struct{}{}
// 	}
// 	for {
// 		if _, ok := tbl[min]; ok {
// 			min++
// 		} else {
// 			break
// 		}
// 	}
// 	return min
// }

// // 暴力解法
// func firstMissingPositive(nums []int) int {
// 	sort.Ints(nums)
// 	min := 1
// 	for i := 0; i < len(nums); i++ {
// 		if min > nums[i] {
// 			continue
// 		} else if min == nums[i] {
// 			min++
// 		} else if min < nums[i] {
// 			break
// 		}
// 	}
// 	return min
// }
