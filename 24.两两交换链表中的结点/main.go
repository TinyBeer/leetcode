package main

import "strconv"

/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := new(ListNode)
	head.Val = nums[0]
	tmp := head
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{Val: nums[i]}
		tmp = tmp.Next
	}
	return head
}

func (l *ListNode) Show() string {
	res := "["
	tmp := l
	if tmp != nil {
		res += strconv.Itoa(tmp.Val)
		tmp = tmp.Next
	}
	for tmp != nil {
		res += "," + strconv.Itoa(tmp.Val)
		tmp = tmp.Next
	}
	res += "]"
	return res
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre := &ListNode{Next: head}
	res := pre

	for pre.Next != nil && pre.Next.Next != nil {
		next := pre.Next
		pre.Next = next.Next
		next.Next = pre.Next.Next
		pre.Next.Next = next
		pre = next
	}
	return res.Next
}
