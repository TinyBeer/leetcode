package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。
请在不改变其结构的情况下，恢复这棵树 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOderTraversal(trees []*TreeNode) string {
	res := make([]string, 0, len(trees))
	if len(trees) == 0 {
		return fmt.Sprintf("%v", res)
	}
	var df func(*TreeNode) []string
	df = func(root *TreeNode) []string {
		if root == nil {
			return []string{"null"}
		}
		res := []string{strconv.Itoa(root.Val)}
		res = append(res, df(root.Left)...)
		res = append(res, df(root.Right)...)
		return res
	}
	for _, root := range trees {
		res = append(res, strings.Join(df(root), ","))
	}

	return fmt.Sprintf("%v", res)
}

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	var traversal func(*TreeNode)
	var pre *TreeNode
	var tmp []*TreeNode
	traversal = func(tn *TreeNode) {
		if tn == nil || len(tmp) == 4 {
			return
		}

		traversal(tn.Left)

		if pre != nil && pre.Val > tn.Val {
			tmp = append(tmp, pre, tn)
			if len(tmp) == 4 {
				return
			}
		}
		pre = tn

		traversal(tn.Right)

	}
	traversal(root)

	fmt.Println(tmp)
	if len(tmp) == 2 {
		tmp[0].Val, tmp[1].Val = tmp[1].Val, tmp[0].Val
	} else {
		tmp[0].Val, tmp[3].Val = tmp[3].Val, tmp[0].Val
	}
}
