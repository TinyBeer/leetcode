package main

/*
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给你一个整数，将其转为罗马数字。
*/

func intToRoman(num int) string {
	var thousands = []string{"", "M", "MM", "MMM"}
	var handreds = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	var tens = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	var ones = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return thousands[num/1000] + handreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}

// // 合并
// func intToRoman(num int) string {
// 	tbl := []struct {
// 		value  int
// 		symbol string
// 	}{
// 		{1000, "M"},
// 		{900, "CM"},
// 		{500, "D"},
// 		{400, "CD"},
// 		{100, "C"},
// 		{90, "XC"},
// 		{50, "L"},
// 		{40, "XL"},
// 		{10, "X"},
// 		{9, "IX"},
// 		{5, "V"},
// 		{4, "IV"},
// 		{1, "I"},
// 	}
// 	res := ""
// 	for _, v := range tbl {
// 		for num >= v.value {
// 			res += v.symbol
// 			num -= v.value
// 		}
// 	}
// 	return res
// }

// // 暴力
// func intToRoman(num int) string {
// 	res := ""
// 	for i := 0; i < num/1000; i++ {
// 		res += "M"
// 	}
// 	num %= 1000
// 	if (num+100)/1000 != 0 {
// 		res += "CM"
// 		num %= 100
// 	}
// 	if num/500 != 0 {
// 		num %= 500
// 		res += "D"
// 		for i := 0; i < num/100; i++ {
// 			res += "C"
// 		}
// 		num %= 100
// 	}

// 	if (num+100)/500 != 0 {
// 		res += "CD"
// 		num %= 100
// 	}

// 	for i := 0; i < num/100; i++ {
// 		res += "C"
// 	}
// 	num %= 100

// 	if (num+10)/100 != 0 {
// 		res += "XC"
// 		num %= 10
// 	}

// 	if num/50 != 0 {
// 		num %= 50
// 		res += "L"
// 		for i := 0; i < num/10; i++ {
// 			res += "X"
// 		}
// 		num %= 10
// 	}

// 	if (num+10)/50 != 0 {
// 		res += "XL"
// 		num %= 10
// 	}

// 	for i := 0; i < num/10; i++ {
// 		res += "X"
// 	}
// 	num %= 10

// 	if (num+1)/10 != 0 {
// 		res += "IX"
// 		return res
// 	}
// 	if (num / 5) != 0 {
// 		num %= 5
// 		res += "V"
// 		for i := 0; i < num; i++ {
// 			res += "I"
// 		}
// 		return res
// 	}

// 	if (num+1)/5 != 0 {
// 		res += "IV"
// 		return res
// 	}
// 	for i := 0; i < num; i++ {
// 		res += "I"
// 	}
// 	return res
// }
