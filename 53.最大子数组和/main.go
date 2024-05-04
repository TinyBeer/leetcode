package main

/*

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组
是数组中的一个连续部分。
*/

func maxSubArray(nums []int) int {
	ans := nums[0]
	tmp := 0
	for _, v := range nums {
		tmp += v
		ans = max(ans, tmp)
		if tmp < 0 {
			tmp = 0
		}
	}
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
