package main

/*
给定两个字符串 s 和 t ，它们只包含小写字母。

字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。
*/

// func findTheDifference(s string, t string) byte {
// 	var tbl [26]int
// 	for i := range s {
// 		tbl[s[i]-'a']++
// 	}

// 	for i := range t {
// 		if tbl[t[i]-'a'] == 0 {
// 			return t[i]
// 		}
// 		tbl[t[i]-'a']--
// 	}
// 	return 0
// }

// func findTheDifference(s, t string) byte {
// 	sum := 0
// 	for _, ch := range s {
// 		sum -= int(ch)
// 	}
// 	for _, ch := range t {
// 		sum += int(ch)
// 	}
// 	return byte(sum)
// }

func findTheDifference(s, t string) (diff byte) {
	for i := range s {
		diff ^= s[i] ^ t[i]
	}
	return diff ^ t[len(t)-1]
}
