package main

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

// 请你将两个数相加，并以相同形式返回一个表示和的链表。

// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1

	helper := 0
	for {
		helper += l1.Val + l2.Val
		l1.Val = helper % 10
		helper /= 10

		if l1.Next == nil {
			l1.Next = l2.Next
			break
		}
		if l2.Next == nil {
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1.Next != nil {
		helper += l1.Next.Val
		l1.Next.Val = helper % 10
		helper /= 10
		l1 = l1.Next
	}

	if helper != 0 {
		l1.Next = &ListNode{Val: helper}
	}

	return head
}
