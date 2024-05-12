package main

import (
	"strings"
)

/*
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，
但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，
返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
*/

func restoreIpAddresses(s string) []string {
	var res []string
	var tmp []string
	var df func(idx int)
	df = func(idx int) {
		if idx == len(s) && len(tmp) == 4 {
			res = append(res, strings.Join(tmp, "."))
			return
		}
		rem := 4 - len(tmp)
		if len(s[idx:]) < rem || len(s[idx:]) > rem*3 {
			return
		}
		for i := 1; i <= 3 && idx+i <= len(s); i++ {
			if check(s[idx : idx+i]) {
				tmp = append(tmp, s[idx:idx+i])
				df(idx + i)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}

	df(0)
	return res
}

func check(s string) bool {
	if len(s) == 1 {
		return true
	} else if s[0] == '0' {
		return false
	}

	var n rune
	for _, c := range s {
		n *= 10
		n += c - '0'
	}
	return n <= 255
}
