package main

import (
	"sort"
)

/*
给你一份航线列表 tickets ，其中 tickets[i] = [fromi, toi] 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序。

所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。

例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。
假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。
*/

func findItinerary(tickets [][]string) []string {
	tbl := make(map[string][]string)
	for _, t := range tickets {
		tbl[t[0]] = append(tbl[t[0]], t[1])
	}
	for _, v := range tbl {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}

	var res []string
	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			if v, ok := tbl[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := tbl[curr][0]
			tbl[curr] = tbl[curr][1:]
			dfs(tmp)
		}
		res = append(res, curr)
	}

	dfs("JFK")
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
