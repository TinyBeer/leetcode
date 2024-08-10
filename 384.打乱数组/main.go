package main

import "math/rand"

/*
给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。打乱后，数组的所有排列应该是 等可能 的。

实现 Solution class:

Solution(int[] nums) 使用整数数组 nums 初始化对象
int[] reset() 重设数组到它的初始状态并返回
int[] shuffle() 返回数组随机打乱后的结果
*/

type Solution struct {
	origin  []int
	shuffle []int
}

func Constructor(nums []int) Solution {
	s := Solution{
		origin: nums,
	}
	s.shuffle = make([]int, len(nums))
	copy(s.shuffle, nums)
	return s
}

func (s *Solution) Reset() []int {
	copy(s.shuffle, s.origin)
	return s.shuffle
}

func (s *Solution) Shuffle() []int {
	rand.Shuffle(len(s.shuffle), func(i, j int) {
		s.shuffle[i], s.shuffle[j] = s.shuffle[j], s.shuffle[i]
	})
	return s.shuffle
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
