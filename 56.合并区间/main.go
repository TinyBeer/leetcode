package main

import "sort"

/*

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，
并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	var cur int = 0
	for i := 1; i < len(intervals); i++ {
		if intervals[cur][1] >= intervals[i][0] {
			intervals[cur][1] = max(intervals[cur][1], intervals[i][1])
		} else {
			res = append(res, intervals[cur])
			cur = i
		}
	}
	res = append(res, intervals[cur])
	return res
}
