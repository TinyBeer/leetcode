package main

import "fmt"

/*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}
	tmp := l1
	for l1.Next != nil {
		if l1.Next.Val < l2.Val {
			l1 = l1.Next
			continue
		}
		helper := l1.Next
		l1.Next = l2
		l2 = l2.Next
		l1 = l1.Next
		l1.Next = helper
		if l2 == nil {
			break
		}
	}
	if l1.Next == nil {
		l1.Next = l2
	}
	return tmp
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	f := &ListNode{Next: head}
	s := f
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}

	f = s.Next
	s.Next = nil
	fmt.Println(s.Val)
	before := sortList(head)
	after := sortList(f)

	return merge(before, after)
}
