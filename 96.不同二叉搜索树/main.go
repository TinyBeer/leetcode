package main

/*
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的
二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
*/

func numTrees(n int) int {
	df := make([]int, n+1)
	df[0] = 1
	df[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			df[i] += df[j-1] * df[i-j]
		}
	}
	return df[n]
}

// // 探索
// func numTrees(n int) int {
// 	return df(1, n)
// }

// func df(l, r int) int {
// 	switch r - l {
// 	case -1, 0:
// 		return 1
// 	}
// 	ans := 0
// 	for i := l; i <= r; i++ {
// 		ans += df(l, i-1) * df(i+1, r)
// 	}
// 	return ans
// }
