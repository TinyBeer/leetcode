package main

import (
	"fmt"
	"sort"
)

/*
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
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
		return head
	}
	tmp := head
	for tmp.Next != nil {
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return head
}
