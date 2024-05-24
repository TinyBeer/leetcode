package main

import (
	"sort"
)

/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:

MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。
*/

type MinStack struct {
	top     int
	arr     []int
	sortArr []int
}

func Constructor() MinStack {
	return MinStack{top: -1}
}

func (ms *MinStack) Push(val int) {
	ms.arr = append(ms.arr, val)
	ms.top++
	idx := sort.Search(len(ms.sortArr), func(i int) bool { return ms.sortArr[i] > val })
	ms.sortArr = append(ms.sortArr[:idx], append([]int{val}, ms.sortArr[idx:]...)...)
}

func (ms *MinStack) Pop() {
	x := ms.arr[ms.top]
	idx := sort.Search(len(ms.sortArr), func(i int) bool { return ms.sortArr[i] >= x })
	ms.sortArr = append(ms.sortArr[:idx], ms.sortArr[idx+1:]...)
	ms.arr = ms.arr[:ms.top]
	ms.top--
}

func (ms *MinStack) Top() int {
	return ms.arr[ms.top]
}

func (ms *MinStack) GetMin() int {
	return ms.sortArr[0]
}
