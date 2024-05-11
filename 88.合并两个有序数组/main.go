package main

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，
其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	i1, i2, p := m-1, n-1, m+n-1
	for i1 >= 0 && i2 >= 0 {
		if nums1[i1] >= nums2[i2] {
			nums1[p] = nums1[i1]
			i1--
		} else {
			nums1[p] = nums2[i2]
			i2--
		}
		p--
	}

	for i2 >= 0 {
		nums1[p] = nums2[i2]
		p--
		i2--
	}

}
