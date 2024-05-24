package main

/*
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。
返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
*/

func reverseWords(s string) string {
	s1 := []byte(s)
	slow := 0
	// 去除多于空格
	for fast := 0; fast < len(s1); fast++ {
		if s1[fast] != ' ' {
			if slow != 0 {
				s1[slow] = ' '
				slow++
			}
			for fast < len(s1) && s1[fast] != ' ' {
				s1[slow] = s1[fast]
				slow++
				fast++
			}
		}
	}
	s1 = s1[0:slow]
	reverse(s1)
	// 反转每个单词
	slow = 0
	for fast := 0; fast <= len(s1); fast++ {
		if fast == len(s1) || s1[fast] == ' ' {
			reverse(s1[slow:fast])
			slow = fast + 1
		}
	}
	return string(s1)
}

func reverse(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// func reverseWords(s string) string {
// 	s = strings.TrimSpace(s)
// 	reg := regexp.MustCompile(`\s+`)
// 	sl := strings.Split(reg.ReplaceAllString(s, " "), " ")
// 	for i, j := 0, len(sl)-1; i < j; {
// 		sl[i], sl[j] = sl[j], sl[i]
// 		i++
// 		j--
// 	}
// 	return strings.Join(sl, " ")
// }
