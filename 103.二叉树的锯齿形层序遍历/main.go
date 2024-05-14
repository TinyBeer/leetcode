package main

/*
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。
（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ress := [][]int{}
	res := []int{}
	count := 0
	treeQ, treeNextQ := []*TreeNode{}, []*TreeNode{}
	treeQ = append(treeQ, root)

	for len(treeQ) != 0 {
		node := treeQ[0]
		treeQ = treeQ[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			treeNextQ = append(treeNextQ, node.Left)
		}
		if node.Right != nil {
			treeNextQ = append(treeNextQ, node.Right)
		}

		if len(treeQ) == 0 {
			if count%2 == 0 {
				ress = append(ress, res)
			} else {
				ress = append(ress, myReverse(res))
			}
			count++
			res = []int{}
			treeQ = append(treeQ, treeNextQ...)
			treeNextQ = []*TreeNode{}
		}
	}

	return ress
}

func myReverse(res []int) []int {
	new := []int{}
	for i := len(res) - 1; i >= 0; i-- {
		new = append(new, res[i])
	}
	return new
}

// func zigzagLevelOrder(root *TreeNode) [][]int {
// 	if root == nil {
// 		return nil
// 	}

// 	q := []*TreeNode{root}
// 	var res [][]int
// 	for len(q) != 0 {
// 		size := len(q)
// 		level := make([]int, 0, size)
// 		if len(res)&1 == 0 {
// 			for i := 0; i < size; i++ {
// 				node := q[0]
// 				q = q[1:]
// 				level = append(level, node.Val)
// 				if node.Left != nil {
// 					q = append(q, node.Left)
// 				}
// 				if node.Right != nil {
// 					q = append(q, node.Right)
// 				}
// 			}
// 		} else {
// 			rq := q
// 			q = make([]*TreeNode, 0)
// 			for i := size - 1; i >= 0; i-- {
// 				node := rq[i]
// 				level = append(level, node.Val)
// 				if node.Right != nil {
// 					q = append([]*TreeNode{node.Right}, q...)
// 				}
// 				if node.Left != nil {
// 					q = append([]*TreeNode{node.Left}, q...)
// 				}
// 			}
// 		}
// 		res = append(res, level)
// 	}
// 	return res
// }
