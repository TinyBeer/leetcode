package main

/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
*/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{binarySearchLeft(nums, 0, len(nums)-1, target), binarySearchRightt(nums, 0, len(nums)-1, target)}
}

func binarySearchLeft(nums []int, l int, r int, target int) int {
	mid := (l + r) / 2
	if nums[mid] == target {
		if mid == 0 || nums[mid-1] != target {
			return mid
		} else {
			return binarySearchLeft(nums, l, mid-1, target)
		}
	}
	if l >= r {
		return -1
	}
	if nums[mid] > target {
		return binarySearchLeft(nums, l, mid-1, target)
	}
	return binarySearchLeft(nums, mid+1, r, target)
}

func binarySearchRightt(nums []int, l int, r int, target int) int {
	mid := (l + r) / 2
	if nums[mid] == target {
		if mid == len(nums)-1 || nums[mid+1] != target {
			return mid
		} else {
			return binarySearchRightt(nums, mid+1, r, target)
		}
	}
	if l >= r {
		return -1
	}
	if nums[mid] > target {
		return binarySearchRightt(nums, l, mid-1, target)
	}
	return binarySearchRightt(nums, mid+1, r, target)
}

// func searchRange(nums []int, target int) []int {
//     leftmost := sort.SearchInts(nums, target)
//     if leftmost == len(nums) || nums[leftmost] != target {
//         return []int{-1, -1}
//     }
//     rightmost := sort.SearchInts(nums, target + 1) - 1
//     return []int{leftmost, rightmost}
// }
