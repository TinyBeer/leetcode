package main

/*
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。
 （即逐层地，从左到右访问所有节点）。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var queue []*TreeNode
	queue = append(queue, root)
	res := make([][]int, 0)
	for len(queue) != 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}

	return res
}

// func levelOrder(root *TreeNode) [][]int {
// 	if root == nil {
// 		return nil
// 	}

// 	q := []*TreeNode{root}
// 	st := 0
// 	var tmp []int
// 	var res [][]int
// 	for i := 0; i < len(q); i++ {
// 		if i == st {
// 			if len(tmp) != 0 {
// 				res = append(res, tmp)
// 			}
// 			tmp = make([]int, 0, len(q)-i)
// 			st = len(q)
// 		}
// 		tmp = append(tmp, q[i].Val)
// 		if q[i].Left != nil {
// 			q = append(q, q[i].Left)
// 		}
// 		if q[i].Right != nil {
// 			q = append(q, q[i].Right)
// 		}
// 	}
// 	if len(tmp) != 0 {
// 		res = append(res, tmp)
// 	}
// 	return res
// }
