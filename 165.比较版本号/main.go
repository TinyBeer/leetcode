package main

/*
给你两个 版本号字符串 version1 和 version2 ，请你比较它们。版本号由被点 '.' 分开的修订号组成。修订号的值 是它 转换为整数 并忽略前导零。

比较版本号时，请按 从左到右的顺序 依次比较它们的修订号。如果其中一个版本字符串的修订号较少，则将缺失的修订号视为 0。

返回规则如下：

如果 version1 < version2 返回 -1，
如果 version1 > version2 返回 1，
除此之外返回 0。
*/

func compareVersion(version1 string, version2 string) int {
	var i, j int
	for i != len(version1) || j != len(version2) {
		var v1, v2 int
		for ; i < len(version1); i++ {
			if version1[i] == '.' {
				i++
				break
			}
			v1 *= 10
			v1 += int(version1[i] - '0')
		}
		for ; j < len(version2); j++ {
			if version2[j] == '.' {
				j++
				break
			}
			v2 *= 10
			v2 += int(version2[j] - '0')
		}
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}

	}
	return 0
}
