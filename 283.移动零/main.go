package main

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。
*/

func moveZeroes(nums []int) {
	x := 0
	for _, n := range nums {
		if n != 0 {
			nums[x] = n
			x++
		}
	}
	for i := x; i < len(nums); i++ {
		nums[i] = 0
	}
}

// func moveZeroes(nums []int) {
// 	i, j := 0, len(nums)
// 	for i < j {
// 		for ; i < j && nums[i] != 0; i++ {
// 		}
// 		x := i
// 		for y := x + 1; y < j; y++ {
// 			if nums[y] != 0 {
// 				nums[x] = nums[y]
// 				x = y
// 			}
// 		}
// 		if x == i {
// 			return
// 		}
// 		nums[x] = 0
// 		i = i + 1
// 		j = x
// 	}
// }
