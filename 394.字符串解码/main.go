package main

/*
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
*/

func decodeString(s string) string {
	var ans []byte
	for i := 0; i < len(s); i++ {
		if s[i] <= '9' && s[i] > 0 {
			repeat := 0
			for s[i] != '[' {
				repeat *= 10
				repeat += int(s[i] - '0')
				i++
			}
			cnt := 1
			j := i + 1
			for {
				if s[j] == '[' {
					cnt++
				} else if s[j] == ']' {
					cnt--
				}
				if cnt == 0 {
					break
				}
				j++
			}
			repeatStr := decodeString(s[i+1 : j])
			for x := 0; x < repeat; x++ {
				ans = append(ans, []byte(repeatStr)...)
			}
			i = j
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
