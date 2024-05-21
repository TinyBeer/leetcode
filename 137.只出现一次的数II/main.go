package main

/*
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题。
*/
// 位运算
func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)

}

// func singleNumber(nums []int) int {
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}
// 	sort.Ints(nums)
// 	p := 0
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] == nums[i-1] {
// 			continue
// 		}
// 		if i-p > 1 {
// 			p = i
// 		} else {
// 			return nums[p]
// 		}
// 	}
// 	return nums[len(nums)-1]
// }
