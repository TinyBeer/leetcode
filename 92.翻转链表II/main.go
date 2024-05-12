package main

import "fmt"

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
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

//  1 2 3 4  5  2 4
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}

	help := &ListNode{Next: head}
	l, r := help, head
	for i := 1; i < right; i++ {
		if i < left {
			l = l.Next
		}
		r = r.Next
	}
	tmp := r.Next
	r.Next = nil
	r = tmp

	rev := l.Next
	for {
		tmp = rev.Next
		rev.Next = r
		r = rev
		if tmp == nil {
			break
		}
		rev = tmp
	}
	l.Next = rev
	return help.Next
}
