package main

import "strings"

/*
将非负整数 num 转换为其对应的英文表示。
*/

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	tbl := []string{
		"Zero",
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Eleven",
		"Twelve",
		"Thirteen",
		"Fourteen",
		"Fifteen",
		"Sixteen",
		"Seventeen",
		"Eighteen",
		"Nineteen",
	}
	tbl2 := map[int]string{
		2: "Twenty",
		3: "Thirty",
		4: "Forty",
		5: "Fifty",
		6: "Sixty",
		7: "Seventy",
		8: "Eighty",
		9: "Ninety",
	}
	// tbl3 := []string{"Hundred", "Thousand", "Million", "Billion"}

	var tmp []string
	n := num / 1_000_000_000
	if n > 0 {
		if n >= 20 {
			tmp = append(tmp, tbl2[n/10])
			if n%10 > 0 {
				tmp = append(tmp, tbl[n%10])
			}
		} else if n > 0 {
			tmp = append(tmp, tbl[n])
		}
		tmp = append(tmp, "Billion")
	}
	num = num % 1_000_000_000

	n = num / 1_000_000
	if n > 0 {
		if n > 0 {
			if n >= 100 {
				tmp = append(tmp, tbl[n/100])
				tmp = append(tmp, "Hundred")
				n %= 100
			}
			if n >= 20 {
				tmp = append(tmp, tbl2[n/10])
				if n%10 > 0 {
					tmp = append(tmp, tbl[n%10])
				}
			} else if n > 0 {
				tmp = append(tmp, tbl[n])
			}
			tmp = append(tmp, "Million")
		}
	}
	num %= 1_000_000

	n = num / 1_000
	if n > 0 {
		if n > 0 {
			if n >= 100 {
				tmp = append(tmp, tbl[n/100])
				tmp = append(tmp, "Hundred")
				n %= 100
			}
			if n >= 20 {
				tmp = append(tmp, tbl2[n/10])
				if n%10 > 0 {
					tmp = append(tmp, tbl[n%10])
				}
			} else if n > 0 {
				tmp = append(tmp, tbl[n])
			}
			tmp = append(tmp, "Thousand")
		}
	}
	num %= 1_000

	n = num / 100
	if n > 0 {
		tmp = append(tmp, tbl[n])
		tmp = append(tmp, "Hundred")
	}
	num %= 100

	if num >= 20 {
		tmp = append(tmp, tbl2[num/10])
		num %= 10
	}
	if num > 0 {
		tmp = append(tmp, tbl[num])
	}

	return strings.Join(tmp, " ")
}
