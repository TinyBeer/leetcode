package main

/*
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

不允许修改 链表。
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

// 快慢针
func detectCycle(head *ListNode) *ListNode {
	f, s := head, head
	for f != nil {
		if f.Next == nil {
			return nil
		}
		f = f.Next.Next
		s = s.Next
		if f == s {
			p := head
			for p != s {
				p = p.Next
				s = s.Next
			}
			return p
		}
	}
	return nil
}

// // hash
// func detectCycle(head *ListNode) *ListNode {
// 	tbl := make(map[*ListNode]struct{})
// 	for head != nil {
// 		if _, ok := tbl[head]; ok {
// 			return head
// 		}
// 		tbl[head] = struct{}{}
// 		head = head.Next
// 	}
// 	return nil
// }
