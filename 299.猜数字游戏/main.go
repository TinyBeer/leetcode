package main

import "fmt"

/*
你在和朋友一起玩 猜数字（Bulls and Cows）游戏，该游戏规则如下：

写出一个秘密数字，并请朋友猜这个数字是多少。朋友每猜测一次，你就会给他一个包含下述信息的提示：

猜测数字中有多少位属于数字和确切位置都猜对了（称为 "Bulls"，公牛），
有多少位属于数字猜对了但是位置不对（称为 "Cows"，奶牛）。也就是说，这次猜测中有多少位非公牛数字可以通过重新排列转换成公牛数字。
给你一个秘密数字 secret 和朋友猜测的数字 guess ，请你返回对朋友这次猜测的提示。

提示的格式为 "xAyB" ，x 是公牛个数， y 是奶牛个数，A 表示公牛，B 表示奶牛。

请注意秘密数字和朋友猜测的数字都可能含有重复数字。
*/

func getHint(secret string, guess string) string {
	have, need := make([]int, 10), make([]int, 10)
	bulls := 0
	for i := range secret {
		if secret[i] == guess[i] {
			bulls++
		} else {
			have[int(guess[i]-'0')]++
			need[int(secret[i]-'0')]++
		}
	}
	cows := 0
	for i := range have {
		cows += min(have[i], need[i])
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
