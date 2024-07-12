package main

/*
给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。
*/

func countSmaller(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return []int{0}
	}
	stack := make([]int, n)
	stack[n-1] = nums[n-1]
	ans := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		stack[i] = max(stack[i+1], nums[i])
		if nums[i] > stack[i+1] {
			ans[i] = n - i - 1
			continue
		}
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				ans[i] += ans[j]
				break
			} else if nums[i] > nums[j] {
				ans[i]++
			}
		}
	}

	return ans
}
