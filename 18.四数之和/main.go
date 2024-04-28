package main

import (
	"sort"
)

/*
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。
*/

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)

	res := [][]int{}
	for first := 0; first < n-3; first++ {
		if nums[first]+nums[first+1]+nums[first+2]+nums[first+3] > target {
			break
		}
		if nums[first]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}

		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < n-2; second++ {
			if nums[first]+nums[second]+nums[second+1]+nums[second+2] > target {
				break
			}
			if nums[first]+nums[second]+nums[n-2]+nums[n-1] < target {
				continue
			}
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			third, fourth := second+1, n-1
			for third < fourth {
				sum := nums[first] + nums[second] + nums[third] + nums[fourth]
				if sum == target {
					res = append(res, []int{nums[first], nums[second], nums[third], nums[fourth]})
					for third < fourth && nums[third] == nums[third+1] {
						third++
					}
					for third < fourth && nums[fourth] == nums[fourth-1] {
						fourth--
					}
					third++
					fourth--
				} else if sum > target {
					fourth--
				} else {
					third++
				}
			}

		}
	}
	return res
}
