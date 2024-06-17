package main

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，
最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	var routs [][]*TreeNode
// 	var tmp []*TreeNode
// 	var find func(tn *TreeNode)
// 	find = func(tn *TreeNode) {
// 		if len(routs) == 2 {
// 			return
// 		}
// 		if tn == nil {
// 			return
// 		}
// 		tmp = append(tmp, tn)
// 		if tn == p || tn == q {
// 			routs = append(routs, append([]*TreeNode(nil), tmp...))
// 		}

// 		find(tn.Left)
// 		find(tn.Right)
// 		tmp = tmp[:len(tmp)-1]
// 	}
// 	find(root)

// 	i := 0
// 	for len(routs[0]) > i+1 && len(routs[1]) >= i+1 && routs[0][i+1] == routs[1][i+1] {
// 		i++
// 	}
// 	return routs[0][i]
// }
