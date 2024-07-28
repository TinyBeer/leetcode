package main

/*
给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。

元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次
*/

func reverseVowels(s string) string {
	isVowel := func(b byte) bool {
		switch b {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		default:
			return false
		}
	}
	i, j := 0, len(s)-1
	tmp := []byte(s)
	for i < j {
		for i < j && !isVowel(tmp[i]) {
			i++
		}
		for i < j && !isVowel(tmp[j]) {
			j--
		}
		tmp[i], tmp[j] = tmp[j], tmp[i]
		i++
		j--
	}
	return string(tmp)
}
