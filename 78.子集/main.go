package main

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的
子集
（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/
func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for i := 0; i < len(nums); i++ {
		n := len(res)
		for j := 0; j < n; j++ {
			res = append(res, append(append([]int(nil), res[j]...), nums[i]))
		}
	}
	return res
}
