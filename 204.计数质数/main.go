package main

/*
给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
*/

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	primes := []int{}
	notPrime := make([]bool, n)

	for i := 2; i < n; i++ {
		if !notPrime[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p >= n {
				break
			}
			notPrime[i*p] = true
			if i%p == 0 {
				break
			}
		}
	}
	return len(primes)
}

// func countPrimes(n int) int {
// 	if n <= 2 {
// 		return 0
// 	}
// 	tbl := []int{2}
// 	for i := 3; i < n; i++ {
// 		flag := false
// 		for _, v := range tbl {
// 			if i%v == 0 {
// 				flag = true
// 				break
// 			}
// 		}
// 		if !flag {
// 			tbl = append(tbl, i)
// 		}
// 	}
// 	return len(tbl)
// }
