package main

/*
请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。

实现 MyStack 类：

void push(int x) 将元素 x 压入栈顶。
int pop() 移除并返回栈顶元素。
int top() 返回栈顶元素。
boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。
*/

type MyStack struct {
	q1, q2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q2 = append(this.q2, x)
	if len(this.q1) > 0 {
		this.q1 = append(this.q2, this.q1...)
	} else {
		this.q1 = this.q2
	}
	this.q2 = nil
}

func (this *MyStack) Pop() int {
	x := this.q1[0]
	this.q1 = this.q1[1:]
	return x
}

func (this *MyStack) Top() int {
	return this.q1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0
}
