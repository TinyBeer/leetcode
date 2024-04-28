package main

/*
 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	p1 := &ListNode{Next: head}
	p2 := p1
	head = p1
	for i := 0; i < n; i++ {
		if p2.Next != nil {
			p2 = p2.Next
		}
	}
	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p1.Next = p1.Next.Next
	return head.Next
}
