package main

/*
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
*/

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, l int, r int, target int) int {
	mid := (l + r) / 2
	if nums[mid] == target {
		return mid
	}
	if l >= r {
		return -1
	}
	if nums[l] > nums[r] {
		resl := binarySearch(nums, l, mid-1, target)
		if resl != -1 {
			return resl
		}

		return binarySearch(nums, mid+1, r, target)
	} else {
		if nums[mid] > target {
			return binarySearch(nums, l, mid-1, target)
		} else {
			return binarySearch(nums, mid+1, r, target)
		}
	}
}
