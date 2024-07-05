package main

/*

给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。
*/

type NumMatrix struct {
	arr [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])

	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j > 0 {
					arr[i][j] = arr[i][j-1]
				}
			} else {
				if j == 0 {
					arr[i][j] = arr[i-1][j]
				} else {
					arr[i][j] = arr[i-1][j] + arr[i][j-1] - arr[i-1][j-1]
				}
			}
			arr[i][j] += matrix[i][j]
		}
	}
	return NumMatrix{arr: arr}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := this.arr[row2][col2]
	if row1 != 0 {
		ans -= this.arr[row1-1][col2]
	}
	if col1 != 0 {
		ans -= this.arr[row2][col1-1]
	}
	if row1*col1 != 0 {
		ans += this.arr[row1-1][col1-1]
	}
	return ans
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
