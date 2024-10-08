package main

import "math/rand"

/*
实现RandomizedSet 类：

RandomizedSet() 初始化 RandomizedSet 对象
bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。
*/

type RandomizedSet struct {
	tbl map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		tbl: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.tbl[val]; ok {
		return false
	}
	this.tbl[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.tbl[val]; !ok {
		return false
	} else {
		lastVal := this.arr[len(this.arr)-1]
		this.arr[idx] = lastVal
		this.tbl[lastVal] = idx
		this.arr = this.arr[:len(this.arr)-1]
		delete(this.tbl, val)
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
