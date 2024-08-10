package main

import "math/rand"

/*
给你一个单链表，随机选择链表的一个节点，并返回相应的节点值。每个节点 被选中的概率一样 。

实现 Solution 类：

Solution(ListNode head) 使用整数数组初始化对象。
int getRandom() 从链表中随机选择一个节点并返回该节点的值。链表中所有节点被选中的概率相等。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	arr []int
}

func Constructor(head *ListNode) Solution {
	var arr []int
	tmp := head
	for tmp != nil {
		arr = append(arr, tmp.Val)
		tmp = tmp.Next
	}
	return Solution{
		arr: arr,
	}
}

func (this *Solution) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
