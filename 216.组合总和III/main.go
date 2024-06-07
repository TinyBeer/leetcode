package main

/*
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：

只使用数字1到9
每个数字 最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
*/

func combinationSum3(k int, n int) [][]int {
	tmp := make([]int, k)
	min := (1 + k) * k / 2
	max := (18 - k + 1) * k / 2
	if n < min || n > max {
		return nil
	}

	var total int
	var ans [][]int
	var dfs func(int)
	dfs = func(i int) {
		if i == k {
			if total == n {
				ans = append(ans, append([]int(nil), tmp...))
			}
			return
		}
		if total >= n {
			return
		}
		if i == 0 {
			for x := 1; x <= 9; x++ {
				total += x
				tmp[i] = x
				dfs(i + 1)
				total -= x
			}
		} else {
			for x := tmp[i-1] + 1; x <= 9; x++ {
				total += x
				tmp[i] = x
				dfs(i + 1)
				total -= x

			}
		}
	}
	dfs(0)
	return ans
}
