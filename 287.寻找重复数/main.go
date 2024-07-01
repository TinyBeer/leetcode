package main

/*
	给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。

你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
*/

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// func findDuplicate(nums []int) int {
// 	n := len(nums)
// 	ans := 0
// 	bit_max := 31
// 	for ((n - 1) >> bit_max) == 0 {
// 		bit_max--
// 	}
// 	for bit := 0; bit <= bit_max; bit++ {
// 		x, y := 0, 0
// 		for i := 0; i < n; i++ {
// 			if (nums[i] & (1 << bit)) > 0 {
// 				x++
// 			}
// 			if i >= 1 && (i&(1<<bit)) > 0 {
// 				y++
// 			}
// 		}
// 		if x > y {
// 			ans |= 1 << bit
// 		}
// 	}
// 	return ans
// }

// func findDuplicate(nums []int) int {
// 	n := len(nums)
// 	l, r := 0, n-1
// 	ans := -1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		cnt := 0
// 		for i := 0; i < n; i++ {
// 			if nums[i] <= mid {
// 				cnt++
// 			}
// 		}
// 		if cnt <= mid {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 			ans = mid
// 		}
// 	}
// 	return ans
// }
