package main

/*
给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。

另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。

返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。

注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。

注意：未在等式列表中出现的变量是未定义的，因此无法确定它们的答案。
*/

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	parent := make([]int, len(equations)*2)
	weight := make([]float64, len(equations)*2)
	for i := range parent {
		parent[i] = i
		weight[i] = 1
	}
	var find func(x int) int
	find = func(x int) int {
		if x != parent[x] {
			origin := parent[x]
			parent[x] = find(parent[x])
			weight[x] *= weight[origin]
		}
		return parent[x]
	}

	union := func(x, y int, val float64) {
		rootX := find(x)
		rootY := find(y)
		if rootX == rootY {
			return
		}
		parent[rootX] = rootY
		weight[rootX] = weight[y] * val / weight[x]
	}

	isConnected := func(x, y int) float64 {
		rootX := find(x)
		rootY := find(y)
		if rootX == rootY {
			return weight[x] / weight[y]
		}
		return -1
	}

	tbl := make(map[string]int)
	idx := 0
	setIdx := func(key string) {
		if _, ok := tbl[key]; !ok {
			tbl[key] = idx
			idx++
		}
	}
	for i, e := range equations {
		setIdx(e[0])
		setIdx(e[1])
		union(tbl[e[0]], tbl[e[1]], values[i])
	}
	var res []float64
	for _, q := range queries {
		id1, has1 := tbl[q[0]]
		id2, has2 := tbl[q[1]]
		if !has1 || !has2 {
			res = append(res, -1)
			continue
		}
		res = append(res, isConnected(id1, id2))
	}
	return res
}
