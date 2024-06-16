package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	var tmp []int
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	mid := len(tmp) / 2
	l, r := mid, mid
	if len(tmp)&1 == 0 {
		l--
	}
	for l >= 0 {
		if tmp[l] != tmp[r] {
			return false
		}
		l--
		r++
	}
	return true
}
