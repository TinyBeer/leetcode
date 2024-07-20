package main

import "strings"

/*
序列化二叉树的一种方法是使用 前序遍历 。
当我们遇到一个非空节点时，我们可以记录下这个节点的值。
如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
*/

func isValidSerialization(preorder string) bool {
	if preorder == "#" {
		return true
	}
	nodes := strings.Split(preorder, ",")
	if nodes[0] == "#" {
		return false
	}
	idx := 1
	var dfs func() bool
	dfs = func() bool {
		for i := 0; i < 2; i++ {
			if idx >= len(nodes) {
				return false
			}
			if nodes[idx] == "#" {
				idx++
				continue
			}

			idx++
			if !dfs() {
				return false
			}
		}
		return true
	}
	return dfs() && idx == len(nodes)
}
