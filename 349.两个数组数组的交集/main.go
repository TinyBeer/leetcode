package main

/*
给定两个数组 nums1 和 nums2 ，返回 它们的
交集
 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
*/

func intersection(nums1 []int, nums2 []int) []int {
	tbl := make(map[int]struct{})
	for _, num := range nums1 {
		tbl[num] = struct{}{}
	}
	var res []int
	for _, num := range nums2 {
		if _, ok := tbl[num]; ok {
			res = append(res, num)
			delete(tbl, num)
		}
	}
	return res
}
