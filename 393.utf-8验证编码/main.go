package main

/*
给定一个表示数据的整数数组 data ，返回它是否为有效的 UTF-8 编码。

UTF-8 中的一个字符可能的长度为 1 到 4 字节，遵循以下的规则：

对于 1 字节 的字符，字节的第一位设为 0 ，后面 7 位为这个符号的 unicode 码。
对于 n 字节 的字符 (n > 1)，第一个字节的前 n 位都设为1，第 n+1 位设为 0 ，后面字节的前两位一律设为 10 。剩下的没有提及的二进制位，全部为这个符号的 unicode 码。
*/

func validUtf8(data []int) bool {
	n := len(data)
	if n == 0 {
		return true
	}
	l := 0
	switch {
	case data[0]&0x00000080 == 0:
		return validUtf8(data[1:])
	// case data[0]&0x000000ff == 0x000000ff:
	// 	l = 8
	// case data[0]&0x000000fe == 0x000000fe:
	// 	l = 7
	// case data[0]&0x000000fc == 0x000000fc:
	// 	l = 6
	// case data[0]&0x000000f8 == 0x000000f8:
	// 	l = 5
	case data[0]&0x000000f8 == 0x000000f0:
		l = 4
	case data[0]&0x000000f0 == 0x000000e0:
		l = 3
	case data[0]&0x000000e0 == 0x000000c0:
		l = 2
	default:
		return false
	}
	if n < l {
		return false
	}
	for i := 1; i < l; i++ {
		if data[i]&0x000000c0 != 0x00000080 {
			return false
		}
	}
	return validUtf8(data[l:])
}
