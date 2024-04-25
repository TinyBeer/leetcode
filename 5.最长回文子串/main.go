package main

import (
	"strings"
)

/*
给你一个字符串 s，找到 s 中最长的回文
子串
子字符串
子字符串 是字符串中连续的字符序列。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

*/

// manacher算法
func longestPalindrome(s string) string {
	// 统一奇偶长度的字符串
	t := "#"
	for _, r := range s {
		t += string(r) + "#"
	}

	start, end := 0, -1
	arm_len_list := []int{}
	center, arm_len := 0, 0
	for i := range t {
		var cur_arm_len int
		if center+arm_len > i {
			i_sym := 2*center - i
			min_arm_len := min(arm_len_list[i_sym], center+arm_len-i)
			cur_arm_len = expand(t, i-min_arm_len, i+min_arm_len)
		} else {
			cur_arm_len = expand(t, i, i)
		}
		arm_len_list = append(arm_len_list, cur_arm_len)
		if i+cur_arm_len > center+arm_len {
			center, arm_len = i, cur_arm_len
		}
		if cur_arm_len*2+1 > end-start+1 {
			start = i - cur_arm_len
			end = i + cur_arm_len
		}
	}

	return strings.ReplaceAll(string(t[start:end+1]), "#", "")
}

func expand(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}
	return (right - left - 2) / 2
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// func longestPalindrome(s string) string {
// 	maxStr := ""
// 	for c := range s {
// 		for i := 0; c-i >= 0 && c+i < len(s); i++ {
// 			if s[c-i] == s[c+i] {
// 				if len(maxStr) < 2*i+1 {
// 					maxStr = string(s[c-i : c+i+1])
// 				}
// 			} else {
// 				break
// 			}
// 		}
// 		for i := 0; c-i >= 0 && c+i+1 < len(s); i++ {
// 			if s[c-i] == s[c+i+1] {
// 				if len(maxStr) < 2*i+2 {
// 					maxStr = string(s[c-i : c+i+2])
// 				}
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	return maxStr
// }
