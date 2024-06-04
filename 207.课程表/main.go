package main

/*
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		valid   = true
		dfs     func(u int)
	)

	for _, paire := range prerequisites {
		edges[paire[1]] = append(edges[paire[1]], paire[0])
	}
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
	}
	for i := 0; i < numCourses && valid; i++ {
		dfs(i)
	}
	return valid
}

// func canFinish(numCourses int, prerequisites [][]int) bool {
// 	edges := make([][]int, numCourses)
// 	indeg := make([]int, numCourses)
// 	q := make([]int, 0, numCourses)
// 	for _, pair := range prerequisites {
// 		edges[pair[1]] = append(edges[pair[1]], pair[0])
// 		indeg[pair[0]]++
// 	}

// 	for i := 0; i < numCourses; i++ {
// 		if indeg[i] == 0 {
// 			q = append(q, i)
// 		}
// 	}
// 	for i := 0; i < len(q); i++ {
// 		cur := q[i]
// 		for _, v := range edges[cur] {
// 			indeg[v]--
// 			if indeg[v] == 0 {
// 				q = append(q, v)
// 			}
// 		}
// 	}
// 	return len(q) == numCourses
// }
