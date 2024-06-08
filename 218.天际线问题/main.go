package main

import (
	"container/heap"
	"sort"
)

/*
城市的 天际线 是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。给你所有建筑物的位置和高度，请返回 由这些建筑物形成的 天际线 。

每个建筑物的几何信息由数组 buildings 表示，其中三元组 buildings[i] = [lefti, righti, heighti] 表示：

lefti 是第 i 座建筑物左边缘的 x 坐标。
righti 是第 i 座建筑物右边缘的 x 坐标。
heighti 是第 i 座建筑物的高度。
你可以假设所有的建筑都是完美的长方形，在高度为 0 的绝对平坦的表面上。

天际线 应该表示为由 “关键点” 组成的列表，格式 [[x1,y1],[x2,y2],...] ，并按 x 坐标 进行 排序 。关键点是水平线段的左端点。列表中最后一个点是最右侧建筑物的终点，y 坐标始终为 0 ，仅用于标记天际线的终点。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。

注意：输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...] 是不正确的答案；三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]
*/

type (
	pair struct{ right, height int }
	hp   []pair
)

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }

// func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
// 由于pop结果不需要用到，所以这里可以优化一下
func (h *hp) Pop() interface{} { a := *h; *h = a[:len(a)-1]; return 0 }

func getSkyline(buildings [][]int) (ans [][]int) {
	n := len(buildings)
	boundaries := make([]int, 0, n*2)
	for _, building := range buildings {
		boundaries = append(boundaries, building[0], building[1])
	}
	sort.Ints(boundaries)

	idx := 0
	h := hp{}
	for _, boundary := range boundaries {
		for idx < n && buildings[idx][0] <= boundary {
			heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
			idx++
		}
		for len(h) > 0 && h[0].right <= boundary {
			heap.Pop(&h)
		}

		maxn := 0
		if len(h) > 0 {
			maxn = h[0].height
		}
		if len(ans) == 0 || maxn != ans[len(ans)-1][1] {
			ans = append(ans, []int{boundary, maxn})
		}
	}
	return
}

// func getSkyline(buildings [][]int) [][]int {
// 	var height []int
// 	for _, b := range buildings {
// 		x, y, h := b[0], b[1], b[2]
// 		i := 0
// 		for ; len(height) < x; i++ {
// 			height = append(height, 0)
// 		}
// 		for i = x; i < len(height) && i <= y; i++ {
// 			height[i] = max(height[i], h)
// 		}
// 		for ; i <= y; i++ {
// 			height = append(height, h)
// 		}
// 	}

// 	preH := 0
// 	var ans [][]int
// 	for x, h := range height {
// 		if preH > h {
// 			ans = append(ans, []int{x - 1, h})
// 		} else if preH < h {
// 			ans = append(ans, []int{x, h})
// 		}
// 		preH = h
// 	}
// 	if preH != 0 {
// 		ans = append(ans, []int{len(height) - 1, 0})
// 	}
// 	return ans
// }
