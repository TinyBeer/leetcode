package main

/*
给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。

第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。

请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。

你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	eh := &ListNode{}
	et := eh
	ot := head

	for ot.Next != nil {
		et.Next = ot.Next
		et = et.Next
		if et.Next == nil {
			break
		}
		ot.Next = et.Next
		et.Next = nil
		ot = ot.Next
	}
	ot.Next = eh.Next
	return head
}
