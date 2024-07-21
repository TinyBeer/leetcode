package main

/*
给你一个整数数组 distance 。

从 X-Y 平面上的点 (0,0) 开始，先向北移动 distance[0] 米，然后向西移动 distance[1] 米，向南移动 distance[2] 米，向东移动 distance[3] 米，持续移动。也就是说，每次移动后你的方位会发生逆时针变化。

判断你所经过的路径是否相交。如果相交，返回 true ；否则，返回 false 。
*/

func isSelfCrossing(distance []int) bool {
	for i := 3; i < len(distance); i++ {
		// 第 1 类路径交叉的情况
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}

		// 第 2 类路径交叉的情况
		if i == 4 && distance[3] == distance[1] &&
			distance[4] >= distance[2]-distance[0] {
			return true
		}

		// 第 3 类路径交叉的情况
		if i >= 5 && distance[i-3]-distance[i-5] <= distance[i-1] &&
			distance[i-1] <= distance[i-3] &&
			distance[i] >= distance[i-2]-distance[i-4] &&
			distance[i-2] > distance[i-4] {
			return true
		}
	}
	return false
}

// func isSelfCrossing(distance []int) bool {
// 	visit := make(map[[2]int]struct{})
// 	visit[[2]int{0, 0}] = struct{}{}
// 	x, y := 0, 0
// 	for i := 0; i < len(distance); i++ {
// 		switch i % 4 {
// 		case 0:
// 			for j := 0; j < distance[i]; j++ {
// 				y++
// 				if _, ok := visit[[2]int{x, y}]; ok {
// 					return true
// 				}
// 				visit[[2]int{x, y}] = struct{}{}
// 			}
// 		case 1:
// 			for j := 0; j < distance[i]; j++ {
// 				x--
// 				if _, ok := visit[[2]int{x, y}]; ok {
// 					return true
// 				}
// 				visit[[2]int{x, y}] = struct{}{}
// 			}
// 		case 2:
// 			for j := 0; j < distance[i]; j++ {
// 				y--
// 				if _, ok := visit[[2]int{x, y}]; ok {
// 					return true
// 				}
// 				visit[[2]int{x, y}] = struct{}{}
// 			}
// 		case 3:
// 			for j := 0; j < distance[i]; j++ {
// 				x++
// 				if _, ok := visit[[2]int{x, y}]; ok {
// 					return true
// 				}
// 				visit[[2]int{x, y}] = struct{}{}
// 			}
// 		}
// 	}
// 	return false
// }
