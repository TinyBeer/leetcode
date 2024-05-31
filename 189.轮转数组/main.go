package main

/*
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
*/

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

// func rotate(nums []int, k int) {
// 	if len(nums) < 2 || k == 0 {
// 		return
// 	}
// 	k %= len(nums)
// 	n := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
// 	copy(nums, n)
// }
