package main

/*
给你两个整数 a 和 b ，不使用 运算符 + 和 - ​​​​​​​，计算并返回两整数之和。
*/

func getSum(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}

// func getSum(a int, b int) int {
//     ans := 0
//     tmp := 1
//     shift := 0
//     for i:=0;i<64;i++{
//         if a & tmp != 0 {
//             shift++
//         }
//         if b & tmp != 0 {
//             shift++
//         }
//         switch shift {
//             case 1:
//                 ans |= tmp
//                 shift = 0
//             case 2:
//                 shift = 1
//             case 3:
//                 ans |= tmp
//                 shift = 1
//             default:
//         }
//         tmp <<= 1
//     }
//     return ans

// }
