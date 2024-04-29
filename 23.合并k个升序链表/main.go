package main

import (
	"strconv"
)

/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。
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

// 分而治之
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	left := mergeKLists(lists[:l/2])
	right := mergeKLists(lists[l/2:])
	return helper(left, right)
}
func helper(l, r *ListNode) *ListNode {
	dummy := &ListNode{}
	tmp := dummy
	for l != nil && r != nil {
		if l.Val < r.Val {
			tmp.Next = l
			l = l.Next
		} else {
			tmp.Next = r
			r = r.Next
		}
		tmp = tmp.Next
	}
	if l != nil {
		tmp.Next = l
	} else {
		tmp.Next = r
	}
	return dummy.Next
}

// 堆
// func mergeKLists(lists []*ListNode) *ListNode {
// 	if len(lists) == 0 {
// 		return nil
// 	}
// 	if len(lists) == 1 {
// 		return lists[0]
// 	}
// 	h := &Heap{}
// 	for _, v := range lists {
// 		h.Push(v)
// 	}

// 	head := &ListNode{}
// 	help := head
// 	for h.len() != 0 {
// 		x := h.Pop()
// 		help.Next = x
// 		help = help.Next
// 		x = x.Next
// 		h.Push(x)
// 	}

// 	return head.Next
// }

// type Heap struct {
// 	lists []*ListNode
// }

// func (h *Heap) Push(x *ListNode) {
// 	if x == nil {
// 		return
// 	}
// 	h.push(x)
// 	h.up(h.len() - 1)
// }

// func (h *Heap) Pop() *ListNode {
// 	if h.len() == 0 {
// 		return nil
// 	}

// 	h.swap(0, h.len()-1)
// 	x := h.pop()
// 	h.down(0)
// 	return x

// }

// func (h *Heap) up(i0 int) {
// 	for i0 > 0 {
// 		i := (i0 - 1) / 2
// 		if h.less(i, i0) {
// 			break
// 		}
// 		h.swap(i, i0)
// 		i0 = i
// 	}
// }

// func (h *Heap) down(i0 int) bool {
// 	i := i0
// 	for {
// 		j1 := 2*i + 1
// 		if j1 >= h.len() {
// 			break
// 		}
// 		j := j1
// 		if j2 := j1 + 1; j2 < h.len() && h.less(j2, j1) {
// 			j = j2
// 		}
// 		if !h.less(j, i) {
// 			break
// 		}
// 		h.swap(i, j)
// 		i = j
// 	}
// 	return i > i0
// }

// func (h *Heap) push(x *ListNode) {
// 	if x == nil {
// 		return
// 	}
// 	h.lists = append(h.lists, x)
// }

// func (h *Heap) pop() *ListNode {
// 	x := h.lists[len(h.lists)-1]
// 	h.lists = h.lists[:len(h.lists)-1]
// 	return x
// }

// func (h *Heap) len() int {
// 	return len(h.lists)
// }

// func (h *Heap) swap(i int, j int) {
// 	h.lists[i], h.lists[j] = h.lists[j], h.lists[i]
// }

// func (h *Heap) less(i int, j int) bool {
// 	return h.lists[i].Val < h.lists[j].Val
// }
