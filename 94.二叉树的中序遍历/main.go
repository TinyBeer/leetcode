package main

/*
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenBinaryTree(arr []int) *TreeNode {
	var genNode func(int) *TreeNode
	genNode = func(i int) *TreeNode {
		if i >= len(arr) || arr[i] == -1 {
			return nil
		}
		root := &TreeNode{Val: arr[i]}
		root.Left = genNode(i + 1)
		root.Right = genNode(i + 2)

		return root
	}

	return genNode(0)
}

func inorderTraversal(root *TreeNode) []int {
	var res []int

	var df func(root *TreeNode)
	df = func(root *TreeNode) {
		if root == nil {
			return
		}
		df(root.Left)
		res = append(res, root.Val)
		df(root.Right)
	}
	df(root)
	return res
}
