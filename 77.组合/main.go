package main

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。
*/

func combine(n int, k int) [][]int {
	var res [][]int
	var tmp []int
	var df func(int)
	df = func(i int) {
		if len(tmp) == k {
			res = append(res, append([]int(nil), tmp...))
			return
		}

		for ; i <= n; i++ {
			tmp = append(tmp, i)
			df(i + 1)
			tmp = tmp[:len(tmp)-1]
		}

	}
	df(1)

	return res
}
