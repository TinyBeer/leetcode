package main

/*
给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
*/

func majorityElement(nums []int) []int {
	element1, element2 := 0, 0
	vote1, vote2 := 0, 0

	for _, num := range nums {
		if vote1 > 0 && num == element1 { // 如果该元素为第一个元素，则计数加1
			vote1++
		} else if vote2 > 0 && num == element2 { // 如果该元素为第二个元素，则计数加1
			vote2++
		} else if vote1 == 0 { // 选择第一个元素
			element1 = num
			vote1++
		} else if vote2 == 0 { // 选择第二个元素
			element2 = num
			vote2++
		} else { // 如果三个元素均不相同，则相互抵消1次
			vote1--
			vote2--
		}
	}

	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		if vote1 > 0 && num == element1 {
			cnt1++
		}
		if vote2 > 0 && num == element2 {
			cnt2++
		}
	}

	var ans []int
	// 检测元素出现的次数是否满足要求
	if vote1 > 0 && cnt1 > len(nums)/3 {
		ans = append(ans, element1)
	}
	if vote2 > 0 && cnt2 > len(nums)/3 {
		ans = append(ans, element2)
	}
	return ans
}

// func majorityElement(nums []int) []int {
// 	a := len(nums) / 3
// 	tbl := make(map[int]int)
// 	duplicate := make(map[int]bool)
// 	var res []int
// 	for _, n := range nums {
// 		tbl[n]++
// 		if tbl[n] > a && !duplicate[n] {
// 			res = append(res, n)
// 			duplicate[n] = true
// 			if len(res) == 2 {
// 				return res
// 			}
// 		}
// 	}
// 	return res
// }
