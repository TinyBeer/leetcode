package main

/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，
找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var tmp []int
	var df func(*TreeNode, int)
	df = func(tn *TreeNode, target int) {
		// if tn == nil {
		// 	return
		// }
		target -= tn.Val
		tmp = append(tmp, tn.Val)
		if tn.Left == nil && tn.Right == nil && target == 0 {
			res = append(res, append([]int{}, tmp...))
		}
		if tn.Left != nil {
			df(tn.Left, target)
		}
		if tn.Right != nil {
			df(tn.Right, target)
		}
		tmp = tmp[:len(tmp)-1]
	}
	df(root, targetSum)
	return res
}
