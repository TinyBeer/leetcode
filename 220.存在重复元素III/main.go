package main

/*
给你一个整数数组 nums 和两个整数 indexDiff 和 valueDiff 。

找出满足下述条件的下标对 (i, j)：

i != j,
abs(i - j) <= indexDiff
abs(nums[i] - nums[j]) <= valueDiff
如果存在，返回 true ；否则，返回 false 。
*/

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	getId := func(x int, w int) int {
		if x >= 0 {
			return x / w
		}
		return (x+1)/w - 1
	}
	abs := func(x int) int {
		if x >= 0 {
			return x
		}
		return -x
	}

	tbl := make(map[int]int, len(nums))
	for i, x := range nums {
		id := getId(x, valueDiff+1)
		if _, has := tbl[id]; has {
			return true
		}
		if v, has := tbl[id-1]; has && abs(v-x) <= valueDiff {
			return true
		}
		if v, has := tbl[id+1]; has && abs(v-x) <= valueDiff {
			return true
		}
		tbl[id] = x
		if i >= indexDiff {
			delete(tbl, getId(nums[i-indexDiff], valueDiff+1))
		}

	}
	return false
}
