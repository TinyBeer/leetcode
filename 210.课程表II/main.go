package main

/*
现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，
其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。
如果不可能完成所有课程，返回 一个空数组 。
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	edge := make([][]int, numCourses)
	in := make([]int, numCourses)
	q := make([]int, 0, numCourses)
	for _, paire := range prerequisites {
		edge[paire[1]] = append(edge[paire[1]], paire[0])
		in[paire[0]]++
	}
	for i, v := range in {
		if v == 0 {
			q = append(q, i)
		}
	}
	for i := 0; i < len(q); i++ {
		for _, v := range edge[q[i]] {
			in[v]--
			if in[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(q) != numCourses {
		return nil
	}
	return q
}
