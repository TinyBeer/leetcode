package main

func twoSum(nums []int, target int) []int {
	h := make(map[int]int)
	for i, v := range nums {
		j, has := h[target-v]
		if has {
			return []int{j, i}
		}
		h[v] = i
	}
	return nil
}

// func twoSum(nums []int, target int) []int {
// 	for i, v := range nums {
// 		for j, m := range nums[i+1:] {
// 			if v+m == target {
// 				return []int{i, i + j + 1}
// 			}
// 		}
// 	}
// 	return nil
// }
