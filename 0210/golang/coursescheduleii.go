package coursescheduleii

import (
// "fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	adjacencyLists := make([][]int, numCourses)
	visited := make([]int, numCourses)

	for _, prerequisite := range prerequisites {
		course, need := prerequisite[0], prerequisite[1]

		adjacencyLists[course] = append(adjacencyLists[course], need)
		if hasCycle(course, adjacencyLists, visited) {
			return []int{}
		}
	}

	order := []int{}
	for course := 0; course < numCourses; course++ {
		order = trace(course, adjacencyLists, visited, order)
	}

	return order
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

func trace(course int, adjacencyLists [][]int, visited, order []int) []int {
	if visited[course] == 2 {
		return order
	}

	if len(adjacencyLists[course]) == 0 {
		order = append(order, course)
		visited[course] = 2
		return order
	}

	needs := adjacencyLists[course]
	for _, needCourse := range needs {
		order = trace(needCourse, adjacencyLists, visited, order)
	}

	visited[course] = 2
	order = append(order, course)
	return order
}
