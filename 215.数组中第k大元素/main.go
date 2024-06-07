package main

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}

func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

// func findKthLargest(nums []int, k int) int {
// 	heapSize := len(nums)
// 	var down func(int)
// 	down = func(i int) {
// 		l, r, largest := i*2+1, i*2+2, i
// 		if l < heapSize && nums[l] > nums[largest] {
// 			largest = l
// 		}
// 		if r < heapSize && nums[r] > nums[largest] {
// 			largest = r
// 		}
// 		if largest != i {
// 			nums[i], nums[largest] = nums[largest], nums[i]
// 			down(largest)
// 		}
// 	}
// 	for i := heapSize / 2; i >= 0; i-- {
// 		down(i)
// 	}
// 	for i := 0; i < k-1; i++ {
// 		nums[0], nums[heapSize-1] = nums[heapSize-1], nums[0]
// 		heapSize--
// 		down(0)
// 	}

// 	return nums[0]
// }

// func findKthLargest(nums []int, k int) int {
// 	sort.Slice(nums, func(i, j int) bool {
// 		return nums[i] > nums[j]
// 	})
// 	return nums[k-1]
// }
