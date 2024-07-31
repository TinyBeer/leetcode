package main

/*
 给你一个由非负整数 a1, a2, ..., an 组成的数据流输入，请你将到目前为止看到的数字总结为不相交的区间列表。

实现 SummaryRanges 类：

SummaryRanges() 使用一个空数据流初始化对象。
void addNum(int val) 向数据流中加入整数 val 。
int[][] getIntervals() 以不相交区间 [starti, endi] 的列表形式返回对数据流中整数的总结。
*/

type SummaryRanges struct {
	arr [][]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (this *SummaryRanges) AddNum(value int) {
	l, r := 0, len(this.arr)-1
	for l <= r {
		mid := (l + r + 1) / 2
		if this.arr[mid][0] <= value && this.arr[mid][1] >= value {
			return
		}
		if this.arr[mid][0] > value {
			r = mid - 1
		} else if this.arr[mid][1] < value {
			l = mid + 1
		}
	}
	this.arr = append(this.arr[:l], append([][]int{{value, value}}, this.arr[l:]...)...)
	if l+1 < len(this.arr) && this.arr[l+1][0] == value+1 {
		this.arr[l+1][0] = value
		this.arr = append(this.arr[:l], this.arr[l+1:]...)
	}
	if l-1 >= 0 && this.arr[l-1][1] == value-1 {
		this.arr[l-1][1] = this.arr[l][1]
		this.arr = append(this.arr[:l], this.arr[l+1:]...)
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.arr
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */
