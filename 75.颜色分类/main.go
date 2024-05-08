package main

/*
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。
*/

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	partition := func(left int, right int) int {
		index := left
		piv := nums[index]

		for left <= right {
			for left <= right {
				if nums[right] < piv {
					nums[left] = nums[right]
					index = right
					left++
					break
				} else {
					right--
				}
			}
			for left <= right {
				if nums[left] > piv {
					nums[right] = nums[left]
					index = left
					right--
					break
				} else {
					left++
				}
			}
		}
		nums[index] = piv
		return index
	}

	var quick func(int, int)
	quick = func(l, r int) {
		if l < r {
			pivIdx := partition(l, r)
			quick(l, pivIdx-1)
			quick(pivIdx+1, r)
		}
	}
	quick(0, len(nums)-1)

}
