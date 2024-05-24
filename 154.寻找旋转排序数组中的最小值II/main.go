package main

/*
已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。

你必须尽可能减少整个过程的操作步骤。
*/

func findMin(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] { //r始终在右侧上升区间
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		} else { //若二者相等，则最小元素必定在l+1到r之间，因为l和mid相等，故可以去除
			r--
			//r = mid => 错！如：[3,3,1,3], mid==1, r=mid就会越过最小值1
			//nums[mid]==nums[r]不能判断最小值在哪个区间
		}
	}
	return nums[l]
}

// func findMin(nums []int) int {
// 	var df func(int, int) int
// 	df = func(i, j int) int {
// 		if i == j {
// 			return nums[i]
// 		}
// 		if nums[i] < nums[j] {
// 			return nums[i]
// 		}
// 		mid := (i + j) / 2
// 		return min(df(i, mid), df(mid+1, j))
// 	}
// 	return df(0, len(nums)-1)
// }
