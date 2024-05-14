package main

/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var df func(*TreeNode, *TreeNode) bool
	df = func(lt, rt *TreeNode) bool {
		if lt == nil && rt == nil {
			return true
		}
		if (lt == nil && rt != nil) || (lt != nil && rt == nil) || lt.Val != rt.Val {
			return false
		}
		return df(lt.Left, rt.Right) && df(lt.Right, rt.Left)
	}

	return df(root.Left, root.Right)
}
