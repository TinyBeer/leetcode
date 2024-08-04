package main

/*
有两个水壶，容量分别为 x 和 y 升。水的供应是无限的。确定是否有可能使用这两个壶准确得到 target 升。

你可以：

装满任意一个水壶
清空任意一个水壶
将水从一个水壶倒入另一个水壶，直到接水壶已满，或倒水壶已空。
*/
func canMeasureWater(x int, y int, target int) bool {
	if x+y < target {
		return false
	}
	if x*y == 0 {
		return target == 0 || target == x+y
	}
	return target%gcd(x, y) == 0
}

func gcd(x, y int) int {
	remainder := x % y
	for remainder != 0 {
		x = y
		y = remainder
		remainder = x % y
	}
	return y
}
