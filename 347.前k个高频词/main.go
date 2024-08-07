package main

import "container/heap"

/*
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
*/

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// func topKFrequent(nums []int, k int) []int {
// 	if k == 0 {
// 		return nil
// 	}
// 	tbl := make(map[int]int)
// 	for _, v := range nums {
// 		tbl[v]++
// 	}

// 	ans := make([]int, 0, k)
// 	for i := 0; i < k; i++ {
// 		tmp := 0
// 		cnt := 0
// 		for k, v := range tbl {
// 			if v > cnt {
// 				tmp = k
// 				cnt = v
// 			}
// 		}
// 		ans = append(ans, tmp)
// 		tbl[tmp] = 0
// 	}
// 	return ans
// }
