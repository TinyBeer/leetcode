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
	if head == nil || k < 2 {
		return head
	}
	tmp := &ListNode{Next: head}
	pre := tmp
	kpre := tmp
	for {
		for i := 0; i < k; i++ {
			kpre = kpre.Next
			if kpre == nil {
				return tmp.Next
			}
		}

		cur := pre.Next
		next := cur.Next
		link := kpre.Next
		kpre = cur
		preNext := cur
		nextPre := pre.Next
		for i := 0; i < k; i++ {
			preNext = cur
			cur.Next = link
			link = cur
			cur = next
			if cur == nil {
				break
			}
			next = cur.Next
		}
		pre.Next = preNext
		pre = nextPre
	}
}
