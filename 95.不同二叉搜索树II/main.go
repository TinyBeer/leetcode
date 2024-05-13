package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n
互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTree(arrs ...int) *TreeNode {
	if len(arrs) == 0 {
		return nil
	}
	root := &TreeNode{Val: arrs[0]}
	for i := 1; i < len(arrs); i++ {
		insert(root, arrs[i])
	}
	return root
}
func insert(root *TreeNode, val int) {
	if val == -1 {
		return
	}
	if root.Val > val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			insert(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			insert(root.Right, val)
		}
	}
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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return df(1, n)
}

func df(l, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{nil}
	}
	res := make([]*TreeNode, 0, r-l+1)
	for i := l; i <= r; i++ {
		left := df(l, i-1)
		right := df(i+1, r)
		for _, lt := range left {
			for _, rt := range right {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  lt,
					Right: rt,
				})
			}
		}
	}
	return res
}
