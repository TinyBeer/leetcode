package main

import (
	"strconv"
	"strings"
)

/*
给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为
平衡
 二叉搜索树。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenNodeList(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	tmp := head
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{Val: nums[i]}
		tmp = tmp.Next
	}
	return head
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOderTraversal(root *TreeNode) string {
	var df func(*TreeNode) []string
	df = func(tn *TreeNode) []string {
		if tn == nil {
			return []string{"null"}
		}
		return append(append([]string{strconv.Itoa(tn.Val)}, df(tn.Left)...), df(tn.Right)...)
	}
	return strings.Join(df(root), ",")
}

// 暴力
func sortedListToBST(head *ListNode) *TreeNode {
	sortedArr := make([]int, 0)
	for head != nil {
		sortedArr = append(sortedArr, head.Val)
		head = head.Next
	}
	return sortedArrayToBST(sortedArr)
}

func sortedArrayToBST(nums []int) *TreeNode {
	var (
		dfs func(low, high int) *TreeNode
	)
	dfs = func(low, high int) *TreeNode {
		if low > high {
			return nil
		}
		mid := (low + high) / 2
		root := &TreeNode{
			Val: nums[mid],
		}
		root.Left = dfs(low, mid-1)
		root.Right = dfs(mid+1, high)
		return root
	}
	return dfs(0, len(nums)-1)
}
