package main

/*
在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。
你从其中的一个加油站出发，开始时油箱为空。

给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，
否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
*/

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	if n == 0 {
		return -1
	}

	rem, s := 0, 0
	for i := 0; i < n; i++ {
		rem += gas[i] - cost[i]
		if rem < 0 {
			s = i + 1
			rem = 0
		}
	}

	for i := 0; i < s; i++ {
		rem += gas[i] - cost[i]
		if rem < 0 {
			return -1
		}
	}
	return s
}

// // 暴力解法
// func canCompleteCircuit(gas []int, cost []int) int {
// 	for i := 0; i < len(gas); i++ {
// 		flag := true
// 		rem := 0
// 		for j := 0; j < len(gas); j++ {
// 			p := (i + j) % len(gas)
// 			rem += gas[p] - cost[p]
// 			if rem < 0 {
// 				flag = false
// 				break
// 			}
// 		}
// 		if flag {
// 			return i
// 		}
// 	}
// 	return -1
// }
