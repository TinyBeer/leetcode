package main

import "sort"

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/
// 双指针
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

// func threeSum(nums []int) [][]int {
// 	sort.Ints(nums)
// 	ans := [][]int{}
// 	for i := 0; i < len(nums)-2; {
// 		if nums[i] > 0 {
// 			break
// 		}
// 		for j := i + 1; j < len(nums)-1; {
// 			if nums[i]+nums[j]*2 > 0 {
// 				break
// 			}
// 			for k := j + 1; k < len(nums); k++ {
// 				if nums[i]+nums[j]+nums[k] == 0 {
// 					ans = append(ans, []int{nums[i], nums[j], nums[k]})
// 					break
// 				} else if nums[i]+nums[j]+nums[k] > 0 {
// 					break
// 				}
// 			}
// 			cur := nums[j]
// 			for j = j + 1; j < len(nums)-1; j++ {
// 				if cur != nums[j] {
// 					break
// 				}
// 			}
// 		}
// 		cur := nums[i]
// 		for i = i + 1; i < len(nums); i++ {
// 			if cur != nums[i] {
// 				break
// 			}
// 		}
// 	}
// 	return ans
// }
