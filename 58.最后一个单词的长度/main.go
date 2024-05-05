package main

/*
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大
子字符串
。
*/

func lengthOfLastWord(s string) int {
	cnt := 0
	flag := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			flag = true
			cnt++
		}
		if s[i] == ' ' && flag {
			return cnt
		}
	}
	return cnt
}
