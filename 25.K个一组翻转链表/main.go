package main

import (
	"strconv"
)

/*
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := new(ListNode)
	head.Val = nums[0]
	tmp := head
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{Val: nums[i]}
		tmp = tmp.Next
	}
	return head
}

func (l *ListNode) Show() string {
	res := "["
	tmp := l
	if tmp != nil {
		res += strconv.Itoa(tmp.Val)
		tmp = tmp.Next
	}
	for tmp != nil {
		res += "," + strconv.Itoa(tmp.Val)
		tmp = tmp.Next
	}
	res += "]"
	return res
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		// head, tail =
		myReverse(head, tail)
		pre.Next = tail
		pre = head
		head = head.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}
