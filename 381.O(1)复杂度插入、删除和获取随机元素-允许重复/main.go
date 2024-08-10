package main

import "math/rand"

/*
RandomizedCollection 是一种包含数字集合(可能是重复的)的数据结构。它应该支持插入和删除特定元素，以及删除随机元素。

实现 RandomizedCollection 类:

RandomizedCollection()初始化空的 RandomizedCollection 对象。
bool insert(int val) 将一个 val 项插入到集合中，即使该项已经存在。如果该项不存在，则返回 true ，否则返回 false 。
bool remove(int val) 如果存在，从集合中移除一个 val 项。如果该项存在，则返回 true ，否则返回 false 。注意，如果 val 在集合中出现多次，我们只删除其中一个。
int getRandom() 从当前的多个元素集合中返回一个随机元素。每个元素被返回的概率与集合中包含的相同值的数量 线性相关 。
您必须实现类的函数，使每个函数的 平均 时间复杂度为 O(1) 。

注意：生成测试用例时，只有在 RandomizedCollection 中 至少有一项 时，才会调用 getRandom 。
*/

type RandomizedCollection struct {
	idx  map[int]map[int]struct{}
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		idx: map[int]map[int]struct{}{},
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (r *RandomizedCollection) Insert(val int) bool {
	ids, has := r.idx[val]
	if !has {
		ids = map[int]struct{}{}
		r.idx[val] = ids
	}
	ids[len(r.nums)] = struct{}{}
	r.nums = append(r.nums, val)
	return !has
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (r *RandomizedCollection) Remove(val int) bool {
	ids, has := r.idx[val]
	if !has {
		return false
	}
	var i int
	for id := range ids {
		i = id
		break
	}
	n := len(r.nums)
	r.nums[i] = r.nums[n-1]
	delete(ids, i)
	delete(r.idx[r.nums[i]], n-1)
	if i < n-1 {
		r.idx[r.nums[i]][i] = struct{}{}
	}
	if len(ids) == 0 {
		delete(r.idx, val)
	}
	r.nums = r.nums[:n-1]
	return true
}

/** Get a random element from the collection. */
func (r *RandomizedCollection) GetRandom() int {
	return r.nums[rand.Intn(len(r.nums))]
}
