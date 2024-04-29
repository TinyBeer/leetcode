package main

import "strconv"

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	tmp := head
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		tmp = tmp.Next
	}
	return head
}

func (l *ListNode) Show() string {
	res := "["
	tmp := l.Next
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		list1, list2 = list2, list1
	}
	head := list1
	for list2 != nil && list1.Next != nil {
		if list1.Next.Val < list2.Val {
			list1 = list1.Next
		} else {
			tmp := list1.Next
			list1.Next = list2
			list2 = list2.Next
			list1 = list1.Next
			list1.Next = tmp
		}
	}

	if list1.Next == nil {
		list1.Next = list2
	}

	return head
}
