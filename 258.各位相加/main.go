package main

/*
给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。返回这个结果。
*/

func addDigits(num int) int {
	root := num % 9
	if root != 0 {
		return root
	}
	if root == 0 && num > 0 {
		return 9
	}
	return 0
}

// func addDigits(num int) int {
// 	for num >= 10 {
// 		tmp := 0
// 		for num > 0 {
// 			tmp += num % 10
// 			num /= 10
// 		}
// 		num = tmp
// 	}
// 	return num
// }
