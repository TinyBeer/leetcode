package main

/*
给你一个数组 rectangles ，其中 rectangles[i] = [xi, yi, ai, bi] 表示一个坐标轴平行的矩形。这个矩形的左下顶点是 (xi, yi) ，右上顶点是 (ai, bi) 。

如果所有矩形一起精确覆盖了某个矩形区域，则返回 true ；否则，返回 false 。
*/
func isRectangleCover(rectangles [][]int) bool {
	type point struct{ x, y int }
	tbl := make(map[point]int)
	area, minX, minY, maxX, maxY := 0, rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
	for _, rect := range rectangles {
		x, y, a, b := rect[0], rect[1], rect[2], rect[3]
		area += (a - x) * (b - y)

		minX = min(minX, x)
		minY = min(minY, y)
		maxX = max(maxX, a)
		maxY = max(maxY, b)

		tbl[point{x, y}]++
		tbl[point{x, b}]++
		tbl[point{a, y}]++
		tbl[point{a, b}]++
	}

	if area != (maxX-minX)*(maxY-minY) || tbl[point{minX, minY}] != 1 || tbl[point{minX, maxY}] != 1 || tbl[point{maxX, minY}] != 1 || tbl[point{maxX, maxY}] != 1 {
		return false
	}
	delete(tbl, point{minX, minY})
	delete(tbl, point{minX, maxY})
	delete(tbl, point{maxX, minY})
	delete(tbl, point{maxX, maxY})

	for _, c := range tbl {
		if c != 2 && c != 4 {
			return false
		}
	}
	return true
}
