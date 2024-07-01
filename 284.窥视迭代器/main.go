package main

/*
请你在设计一个迭代器，在集成现有迭代器拥有的 hasNext 和 next 操作的基础上，还额外支持 peek 操作。

实现 PeekingIterator 类：

PeekingIterator(Iterator<int> nums) 使用指定整数迭代器 nums 初始化迭代器。
int next() 返回数组中的下一个元素，并将指针移动到下个元素处。
bool hasNext() 如果数组中存在下一个元素，返回 true ；否则，返回 false 。
int peek() 返回数组中的下一个元素，但 不 移动指针。
注意：每种语言可能有不同的构造函数和迭代器 Iterator，但均支持 int next() 和 boolean hasNext() 函数。
*/

type Iterator struct{}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	panic("not implement")
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	panic("not implement")
}

type PeekingIterator struct {
	iter *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter}
}

func (this *PeekingIterator) hasNext() bool {
	return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	tmp := *this.iter
	res := this.iter.next()
	this.iter = &tmp
	return res
}
