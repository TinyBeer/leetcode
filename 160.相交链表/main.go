package main

/*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。
如果两个链表不存在相交节点，返回 null 。

图示两个链表在节点 c1 开始相交：
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	cnt := 0
	for p1 != nil || p2 != nil {
		if p1 != nil {
			p1 = p1.Next
			cnt++
		}
		if p2 != nil {
			p2 = p2.Next
			cnt--
		}
	}
	p1, p2 = headA, headB
	if cnt > 0 {
		for i := 0; i < cnt; i++ {
			p1 = p1.Next
		}
	} else {
		for i := 0; i > cnt; i-- {
			p2 = p2.Next
		}
	}
	for p1 != nil && p2 != nil && p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	if p1 == nil || p2 == nil {
		return nil
	}
	return p1
}
