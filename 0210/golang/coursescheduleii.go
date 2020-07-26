package coursescheduleii

import (
// "fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	adjacencyLists := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		course, need := prerequisite[0], prerequisite[1]

		adjacencyLists[course] = append(adjacencyLists[course], need)
	}

	order := []int{}
	var cycle bool
	visited := make([]bool, numCourses)
	for course := 0; course < numCourses; course++ {
		cycleVisited := make([]bool, numCourses)
		cycle, order = trace(course, adjacencyLists, cycleVisited, visited, order)
		if cycle {
			return []int{}
		}
	}

	return order
}

func trace(course int, adjacencyLists [][]int, thisCycleVisited, visited []bool, order []int) (bool, []int) {
	if thisCycleVisited[course] {
		return true, []int{}
	}

	if visited[course] {
		return false, order
	}

	needs := adjacencyLists[course]
	var cycle bool
	thisCycleVisited[course] = true
	for _, needCourse := range needs {
		cycle, order = trace(needCourse, adjacencyLists, thisCycleVisited, visited, order)
		if cycle {
			return true, order
		}
	}
	thisCycleVisited[course] = false
	visited[course] = true
	order = append(order, course)
	return false, order
}
