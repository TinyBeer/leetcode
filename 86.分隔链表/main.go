package main

import (
	"fmt"
)

/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x
的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{
		Val: nums[0],
	}
	tmp := head
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{
			Val: nums[i],
		}
		tmp = tmp.Next
	}
	return head
}

func (h *ListNode) String() string {
	tmp := h
	var members []int
	for tmp != nil {
		members = append(members, tmp.Val)
		tmp = tmp.Next
	}
	return fmt.Sprintf("%v", members)
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	help := &ListNode{Next: head}
	inPos := help
	for inPos.Next.Val < x {
		if inPos.Next.Next == nil {
			return head
		}
		inPos = inPos.Next
	}
	mov := inPos.Next
	for mov.Next != nil {
		if mov.Next.Val < x {
			sm := mov.Next
			mov.Next = sm.Next
			sm.Next = inPos.Next
			inPos.Next = sm
			inPos = inPos.Next
		} else {
			mov = mov.Next
		}
	}
	return help.Next
}
