package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	tmp := n
	var strs []string
	for tmp != nil {
		strs = append(strs, strconv.Itoa(tmp.Val))
		tmp = tmp.Next
	}
	return fmt.Sprintf("[%s]", strings.Join(strs, ","))
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	fast, slow := head, head
	for i := 0; i < k; i++ {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			k = k % (i + 1)
			if k == 0 {
				return head
			}
			fast = slow
			for j := 0; j < k; j++ {
				fast = fast.Next
			}
			break
		}
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	fast.Next = head
	head = slow.Next
	slow.Next = nil
	return head
}
