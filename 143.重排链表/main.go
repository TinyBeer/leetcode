package main

/*
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	f, s := head, head
	for f != nil {
		f = f.Next
		if f == nil {
			break
		}
		f = f.Next
		s = s.Next
	}
	// 翻转后半
	tmp := s.Next
	s.Next = nil
	s = tmp
	rev := &ListNode{}
	for s != nil {
		tmp = s.Next
		s.Next = rev.Next
		rev.Next = s
		s = tmp
	}
	rev = rev.Next
	// 合并
	h := head
	for rev != nil {
		tmp = h.Next
		h.Next = rev
		s = rev.Next
		rev.Next = tmp
		h = tmp
		rev = s
	}
}
