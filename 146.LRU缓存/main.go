package main

/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；
如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/
type node struct {
	key  int
	val  int
	pre  *node
	next *node
}

type LRUCache struct {
	cnt      int
	capacity int
	header   *node
	tailer   *node
	tbl      map[int]*node
}

func Constructor(capacity int) LRUCache {
	header := &node{}
	tailer := &node{}
	header.next = tailer
	tailer.pre = header
	return LRUCache{
		cnt:      0,
		capacity: capacity,
		header:   header,
		tailer:   tailer,
		tbl:      make(map[int]*node),
	}
}

func (c *LRUCache) Get(key int) int {
	if e, ok := c.tbl[key]; ok {
		pre := e.pre
		next := e.next
		pre.next = next
		next.pre = pre
		e.next = c.header.next
		c.header.next.pre = e
		e.pre = c.header
		c.header.next = e
		return e.val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	e, ok := c.tbl[key]
	if !ok && c.cnt == c.capacity {
		e = c.tailer.pre
		delete(c.tbl, e.key)
		ok = true
	}
	if ok {
		pre := e.pre
		next := e.next
		pre.next = next
		next.pre = pre
	} else {
		e = &node{}
		c.cnt++
	}

	c.tbl[key] = e
	e.key = key
	e.val = value
	e.next = c.header.next
	c.header.next.pre = e
	e.pre = c.header
	c.header.next = e
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
