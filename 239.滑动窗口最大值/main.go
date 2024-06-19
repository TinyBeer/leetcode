package main

import "runtime/debug"

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。
*/

func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	tmp := []int{}
	for i := 0; i < len(nums); i++ {
		for len(tmp) > 0 && nums[i] > nums[tmp[len(tmp)-1]] {
			tmp = tmp[:len(tmp)-1]
		}
		tmp = append(tmp, i)
		if i-tmp[0] == k {
			tmp = tmp[1:]
		}
		if i >= k-1 {
			ans = append(ans, nums[tmp[0]])
		}
	}
	return ans
}

func init() {
	debug.SetGCPercent(-1)
}

// var a []int

// type hp struct{ sort.IntSlice }

// func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
// func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
// func (h *hp) Pop() interface{} {
// 	a := h.IntSlice
// 	v := a[len(a)-1]
// 	h.IntSlice = a[:len(a)-1]
// 	return v
// }

// func maxSlidingWindow(nums []int, k int) []int {
// 	a = nums
// 	q := &hp{make([]int, k)}
// 	for i := 0; i < k; i++ {
// 		q.IntSlice[i] = i
// 	}
// 	heap.Init(q)

// 	n := len(nums)
// 	ans := make([]int, 1, n-k+1)
// 	ans[0] = nums[q.IntSlice[0]]
// 	for i := k; i < n; i++ {
// 		heap.Push(q, i)
// 		for q.IntSlice[0] <= i-k {
// 			heap.Pop(q)
// 		}
// 		ans = append(ans, nums[q.IntSlice[0]])
// 	}
// 	return ans
// }

// func maxSlidingWindow(nums []int, k int) []int {
// 	if k == 1 {
// 		return nums
// 	}

// 	var arr []int
// 	arr = append(arr, nums[:k]...)
// 	sort.Ints(arr)
// 	res := []int{arr[k-1]}

// 	remove := func(x int) {
// 		l, r := 0, k-1
// 		for l <= r {
// 			mid := (l + r) / 2
// 			if arr[mid] == x {
// 				arr = append(arr[:mid], arr[mid+1:]...)
// 				return
// 			} else if arr[mid] > x {
// 				r = mid - 1
// 			} else if arr[mid] < x {
// 				l = mid + 1
// 			}
// 		}
// 	}

// 	insert := func(x int) {
// 		l, r := 0, k-1
// 		for l < r {
// 			mid := (l + r) / 2
// 			if arr[mid] == x {
// 				l = mid
// 				r = mid
// 			} else if arr[mid] > x {
// 				r = mid
// 			} else if arr[mid] < x {
// 				l = mid + 1
// 			}
// 		}
// 		arr = append(arr[:l], append([]int{x}, arr[l:]...)...)
// 	}

// 	for i := k; i < len(nums); i++ {
// 		remove(nums[i-k])
// 		insert(nums[i])
// 		res = append(res, arr[k-1])
// 	}
// 	return res
// }
