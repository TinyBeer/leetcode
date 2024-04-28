package main

import (
	"math"
	"sort"
)

/*
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。
*/
// 双指针
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := math.MaxInt32
	for first := 0; first < n-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		for second := first + 1; second < n-1; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for ; second < third; third-- {
				if third < n-1 && nums[third] == nums[third+1] {
					continue
				}

				sum := nums[first] + nums[second] + nums[third]
				if sum == target {
					return sum
				}
				ans = closest(ans, sum, target)
				if sum < target {
					break
				}
			}
		}
	}
	return ans
}

func closest(x int, y int, target int) int {
	if abs(target-x) < abs(target-y) {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}
