// tags: graphs, star2, medium

// time complexity: O(v+e)
// space complexity: O(v+e)
// `v` is number of courses
// `e` is number of prerequisites
func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		preMap[i] = []int{}
	}
	for _, prerequisite := range prerequisites {
		preMap[prerequisite[0]] = append(preMap[prerequisite[0]], prerequisite[1])
	}

	visited, cycle := make(map[int]bool), make(map[int]bool)
	for i := 0; i < numCourses; i++ {
		if !dfs(i, preMap, visited, cycle) {
			return false
		}
	}

	return true
}

func dfs(course int, preMap map[int][]int, visited, cycle map[int]bool) bool {
	if _, ok := cycle[course]; ok {
		return false
	}
	if _, ok := visited[course]; ok {
		return true
	}
	cycle[course] = true
	for _, preCourse := range preMap[course] {
		if !dfs(preCourse, preMap, visited, cycle) {
			return false
		}
	}
	delete(cycle, course)
	visited[course] = true
	return true
}