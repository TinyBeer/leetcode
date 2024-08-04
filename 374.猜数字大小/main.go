package main

/*
我们正在玩猜数字游戏。猜数字游戏的规则如下：

我会从 1 到 n 随机选择一个数字。 请你猜选出的是哪个数字。

如果你猜错了，我会告诉你，我选出的数字比你猜测的数字大了还是小了。

你可以通过调用一个预先定义好的接口 int guess(int num) 来获取猜测结果，返回值一共有三种可能的情况：

-1：你猜的数字比我选出的数字大 （即 num > pick）。
1：你猜的数字比我选出的数字小 （即 num < pick）。
0：你猜的数字与我选出的数字相等。（即 num == pick）。
返回我选出的数字。
*/
func guess(num int) int

func guessNumber(n int) int {
	l, r := 0, n
	for {
		mid := (l + r) / 2
		switch guess(mid) {
		case 0:
			return mid
		case -1:
			r = mid - 1
		case 1:
			l = mid + 1
		}
	}
}
