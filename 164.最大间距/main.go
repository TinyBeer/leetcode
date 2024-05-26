package main

/*
给定一个无序的数组 nums，返回 数组在排序之后，相邻元素之间最大的差值 。如果数组元素个数小于 2，则返回 0 。

您必须编写一个在「线性时间」内运行并使用「线性额外空间」的算法。
*/

// 基数排序
func maximumGap(nums []int) (ans int) {
	n := len(nums)
	if n < 2 {
		return
	}

	buf := make([]int, n)
	maxVal := max(nums...)
	for exp := 1; exp <= maxVal; exp *= 10 {
		cnt := [10]int{}
		for _, v := range nums {
			digit := v / exp % 10
			cnt[digit]++
		}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			digit := nums[i] / exp % 10
			buf[cnt[digit]-1] = nums[i]
			cnt[digit]--
		}
		copy(nums, buf)
	}

	for i := 1; i < n; i++ {
		ans = max(ans, nums[i]-nums[i-1])
	}
	return
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

// 排序
// func maximumGap(nums []int) int {
// 	if len(nums) < 2 {
// 		return 0
// 	}
// 	sort.Ints(nums)
// 	ans := 0
// 	for i := 1; i < len(nums); i++ {
// 		ans = max(nums[i] - nums[i-1], ans)
// 	}
// 	return ans
// }
