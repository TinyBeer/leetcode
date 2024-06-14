package main

/*
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, k int) []int {
	res := []int{}
	if root == nil {
		return res
	}

	if root.Left != nil {
		res = dfs(root.Left, k)
	}
	res = append(res, root.Val)
	if len(res) >= k {
		return res
	}
	if root.Right != nil {
		res = append(res, dfs(root.Right, k)...)
	}
	return res
}

func kthSmallest(root *TreeNode, k int) int {
	res := dfs(root, k)
	return res[k-1]
}

// func kthSmallest(root *TreeNode, k int) int {
// 	tbl := make(map[*TreeNode]int)

// 	var nodes func(*TreeNode) int
// 	nodes = func(tn *TreeNode) int {
// 		if n, ok := tbl[tn]; ok {
// 			return n
// 		}
// 		if tn == nil {
// 			return 0
// 		}
// 		tbl[tn] = nodes(tn.Left) + nodes(tn.Right) + 1
// 		return tbl[tn]
// 	}

// 	var dfs func(*TreeNode, int) int
// 	dfs = func(tn *TreeNode, k int) int {
// 		l := nodes(tn.Left)
// 		if l == k-1 {
// 			return tn.Val
// 		} else if l < k-1 {
// 			return dfs(tn.Right, k-l-1)
// 		} else {
// 			return dfs(tn.Left, k)
// 		}
// 	}
// 	return dfs(root, k)
// }
