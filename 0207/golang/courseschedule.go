package courseschedule

import (
// "fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	adjacencyLists := make([][]int, numCourses)
	visited := make([]int, numCourses)

	for _, prerequisite := range prerequisites {
		course, need := prerequisite[0], prerequisite[1]
		adjacencyLists[course] = append(adjacencyLists[course], need)

		if hasCycle(course, adjacencyLists, visited) {
			return false
		}
	}
	return true
}

func hasCycle(course int, adjacencyLists [][]int, visited []int) bool {
	if len(adjacencyLists[course]) == 0 {
		return false
	}

	if visited[course] == 2 {
		return true
	}

	visited[course] = 2

	adjacencyList := adjacencyLists[course]
	for _, needCourse := range adjacencyList {
		if hasCycle(needCourse, adjacencyLists, visited) {
			visited[course] = 1
			return true
		}
		visited[course] = 1
	}
	return false
}
