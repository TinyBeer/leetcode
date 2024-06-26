package main

/*
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
*/

func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}

// func missingNumber(nums []int) int {
// 	sort.Ints(nums)
// 	i, j := 0, len(nums)-1
// 	if nums[j] == j {
// 		return j + 1
// 	}

// 	for i < j {
// 		mid := (i + j) / 2
// 		if nums[mid] == mid {
// 			i = mid + 1
// 		} else {
// 			j = mid
// 		}
// 	}
// 	return i
// }
