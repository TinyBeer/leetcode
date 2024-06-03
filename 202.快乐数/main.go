package main

/*
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。
*/

func isHappy(n int) bool {
	tbl := make(map[int]bool)
	pow := [10]int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}

	for {
		tmp := 0
		for n > 0 {
			tmp += pow[n%10]
			n /= 10
		}
		if tmp == 1 {
			return true
		}
		if tbl[tmp] {
			return false
		}
		tbl[tmp] = true
		n = tmp
	}
}
