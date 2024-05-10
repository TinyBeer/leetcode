package main

import (
	"fmt"
	"sort"
)

/*
给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{0, head}

	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
