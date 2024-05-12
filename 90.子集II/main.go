package main

/*
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的
子集
（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
*/

// 暴力
func subsetsWithDup(nums []int) [][]int {
	tbl := make(map[int]int)
	for _, v := range nums {
		tbl[v]++
	}

	res := append([][]int{}, []int{})
	for n, c := range tbl {
		for i := len(res) - 1; i >= 0; i-- {
			tmp := []int{n}
			for j := 0; j < c; j++ {
				res = append(res, append(append([]int{}, res[i]...), tmp...))
				tmp = append(tmp, n)
			}
		}
	}
	return res
}
