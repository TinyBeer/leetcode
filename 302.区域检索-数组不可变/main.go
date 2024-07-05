package main

/*
给定一个整数数组  nums，处理以下类型的多个查询:

计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和 ，其中 left <= right
实现 NumArray 类：

NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int i, int j) 返回数组 nums 中索引 left 和 right 之间的元素的 总和 ，
包含 left 和 right 两点（也就是 nums[left] + nums[left + 1] + ... + nums[right] )
*/

type NumArray struct {
	sumArray []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	na := NumArray{
		sumArray: make([]int, n),
	}
	na.sumArray[0] = nums[0]
	for i := 1; i < n; i++ {
		na.sumArray[i] = na.sumArray[i-1] + nums[i]
	}
	return na
}

func (this *NumArray) SumRange(left int, right int) int {
	ans := this.sumArray[right]
	if left != 0 {
		ans -= this.sumArray[left-1]
	}
	return ans
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
