package main

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	return float64(getK(nums1, 0, m-1, nums2, 0, n-1, (m+n-1)/2)+getK(nums1, 0, m-1, nums2, 0, n-1, (m+n)/2)) / 2
}

func getK(num1 []int, start1 int, end1 int, num2 []int, start2 int, end2 int, k int) int {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1
	if len1 < len2 {
		return getK(num2, start2, end2, num1, start1, end1, k)
	}
	if len2 == 0 {
		return num1[start1+k]
	}
	if k == 0 {
		return min(num1[start1], num2[start2])
	}

	p1, p2 := start1+(k-1)/2, min(end2, start2+(k-1)/2)
	if num1[p1] > num2[p2] {
		return getK(num1, start1, end1, num2, p2+1, end2, k-(p2-start2+1))
	} else {
		return getK(num1, p1+1, end1, num2, start2, end2, k-(p1-start1+1))
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// 1 3 5 8 10 13
// 2 4 6 7 9 14

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	nums1 = append(nums1, nums2...)
// 	sort.Ints(nums1)
// 	mid := len(nums1) >> 1
// 	if len(nums1)&1 == 1 {
// 		return float64(nums1[mid])
// 	}

// 	return float64(nums1[mid-1]+nums1[mid]) / 2
// }
