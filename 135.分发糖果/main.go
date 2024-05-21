package main

/*
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
*/
func candy(ratings []int) int {
	df := make([]int, len(ratings))

	for i := range ratings {
		if i == 0 {
			df[i] = 1
		} else if ratings[i] > ratings[i-1] {
			df[i] = df[i-1] + 1
		} else {
			df[i] = 1
		}
	}

	ans := 0
	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			df[i] = max(df[i], df[i+1]+1)
		}
		ans += df[i]
	}
	return ans
}
