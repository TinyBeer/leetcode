package main

/*
峰值元素是指其值严格大于左右相邻值的元素。

给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，
返回 任何一个峰值 所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞ 。

你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
*/

func findPeakElement(nums []int) int {
	l, r := -1, len(nums)-1
	for l+1 < r {
		mid := (l + r) >> 1
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

// func findPeakElement(nums []int) int {
// 	if len(nums) == 0 {
// 		return -1
// 	}
// 	if len(nums) == 1 {
// 		return 0
// 	}
// 	var df func(int, int) int
// 	df = func(l, r int) int {
// 		mid := (l + r) / 2
// 		if mid == 0 {
// 			if nums[mid] > nums[mid+1] {
// 				return mid
// 			} else {
// 				return df(mid+1, r)
// 			}
// 		}

// 		if mid == len(nums)-1 {
// 			if nums[mid-1] < nums[mid] {
// 				return mid
// 			} else {
// 				return df(l, mid-1)
// 			}
// 		}

// 		if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
// 			return mid
// 		}

// 		if nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1] {
// 			idx := df(mid+1, r)
// 			if idx != -1 {
// 				return idx
// 			}
// 		}
// 		if nums[mid-1] > nums[mid] && nums[mid] > nums[mid+1] {
// 			idx := df(l, mid-1)
// 			if idx != -1 {
// 				return idx
// 			}
// 		}
// 		idx := df(mid+1, r)
// 		if idx != -1 {
// 			return idx
// 		}
// 		idx = df(l, mid-1)
// 		if idx != -1 {
// 			return idx
// 		}
// 		return -1
// 	}
// 	return df(0, len(nums)-1)
// }
