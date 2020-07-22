package courseschedule

import (
// "fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	adjacencyLists := make([][]int, numCourses)

	for _, prerequisite := range prerequisites {
		course, need := prerequisite[0], prerequisite[1]
		adjacencyLists[course] = append(adjacencyLists[course], need)

		if hasCycle(course, adjacencyLists, map[int]bool{}) {
			return false
		}
	}
	return true
}

func hasCycle(course int, adjacencyLists [][]int, visited map[int]bool) bool {
	if len(adjacencyLists[course]) == 0 {
		return false
	}

	if _, found := visited[course]; found {
		return true
	}
	visited[course] = true

	adjacencyList := adjacencyLists[course]
	for _, needCourse := range adjacencyList {
		if hasCycle(needCourse, adjacencyLists, visited) {
			return true
		}
		delete(visited, needCourse)
	}
	return false
}
