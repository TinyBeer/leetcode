package main

import (
	"strings"
)

/*
给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。

你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。

要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

文本的最后一行应为左对齐，且单词之间不插入额外的空格。

注意:

单词是指由非空格字符组成的字符序列。
每个单词的长度大于 0，小于等于 maxWidth。
输入单词数组 words 至少包含一个单词。
*/

func fullJustify(words []string, maxWidth int) []string {
	var ans []string
	var inline []string
	var curLen int

	gen := func() {
		space := maxWidth - curLen
		if len(inline) == 1 {
			ans = append(ans, inline[0]+strings.Repeat(" ", space))
		} else {
			rspace := space / (len(inline) - 1)
			rem := space % (len(inline) - 1)
			line := ""
			for i, word := range inline {
				curspce := rspace
				if rem > 0 {
					curspce += 1
					rem--
				}
				line += word
				if i == len(inline)-1 {
					break
				}
				line += strings.Repeat(" ", curspce)
			}
			ans = append(ans, line)
		}
	}
	for _, word := range words {
		// justicfy
		if curLen+len(inline)+len(word) <= maxWidth {
			inline = append(inline, word)
			curLen += len(word)
			continue
		} else {
			// generate line
			gen()
			inline = []string{word}
			curLen = len(word)
		}
	}
	lastLine := strings.Join(inline, " ")
	ans = append(ans, lastLine+strings.Repeat(" ", maxWidth-len(lastLine)))

	return ans
}
