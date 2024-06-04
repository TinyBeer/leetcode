package main

/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var ptr, tmp *ListNode
	for head != nil {
		tmp = head.Next
		head.Next = ptr
		ptr = head
		head = tmp
	}
	return ptr
}
