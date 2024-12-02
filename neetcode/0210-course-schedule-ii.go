// tags: graphs, star3, topological-sort, medium

// time complexity: O(v+e)
// space complexity: O(v+e)
// `v` is number of courses
// `e` is number of prerequisites
func findOrder(numCourses int, prerequisites [][]int) []int {
	preMap := make(map[int][]int)
	for _, prerequisite := range prerequisites {
		preMap[prerequisite[0]] = append(preMap[prerequisite[0]], prerequisite[1])
	}

	var (
		dfs func(course int) bool
		res []int
	)
	visited, cycle := make(map[int]bool), make(map[int]bool)
	dfs = func(course int) bool {
		if _, ok := cycle[course]; ok {
			return false
		}
		if _, ok := visited[course]; ok {
			return true
		}

		cycle[course] = true
		for _, pre := range preMap[course] {
			if !dfs(pre) {
				return false
			}
		}
		delete(cycle, course)

		res = append(res, course)
		visited[course] = true

		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return []int{}
		}
	}
	return res
}
