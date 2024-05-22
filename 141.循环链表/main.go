package main

/*
给你一个链表的头节点 head ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。

如果链表中存在环 ，则返回 true 。 否则，返回 false 。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func genList(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{}
	help := head
	var tmp *ListNode
	for i := range nums {
		help.Next = &ListNode{Val: nums[i]}
		if i == pos {
			tmp = help.Next
		}
		help = help.Next

	}
	help.Next = tmp
	return head.Next
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	help := head.Next
	for help != nil && help.Next != nil {
		if head == help {
			return true
		}
		head = head.Next
		help = help.Next.Next
	}
	return false
}
