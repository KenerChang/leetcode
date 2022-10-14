package criticalconnectionsinanetwork

import "math"

func dfsUtil(graph [][]int, u int, time int, parent []int, disc []int, low []int, visited []bool) ([][]int, int) {
	visited[u] = true

	disc[u] = time
	low[u] = time

	time = time + 1

	results := [][]int{}
	var bridges [][]int
	for _, v := range graph[u] {
		if !visited[v] {
			parent[v] = u

			bridges, time = dfsUtil(graph, v, time, parent, disc, low, visited)
			results = append(results, bridges...)
			low[u] = min(low[u], low[v])

			if low[v] > disc[u] {
				results = append(results, []int{u, v})
			}

		} else if v != parent[u] {
			low[u] = min(low[u], disc[v])
		}
	}

	return results, time
}

func criticalConnections(n int, connections [][]int) [][]int {
	// This problem is same as finding the bridges in a graph.
	// We can solve this by using Tarjan's algorithm.

	if n == 2 {
		return connections
	}

	// build the graph
	graph := make([][]int, n)
	for _, connection := range connections {
		start, to := connection[0], connection[1]
		graph[start] = append(graph[start], to)
		graph[to] = append(graph[to], start)
	}

	// init utils
	parent := make([]int, n)
	disc := make([]int, n)
	low := make([]int, n)
	visited := make([]bool, n)
	time := 0
	for i := 0; i < n; i++ {
		parent[i] = -1
		disc[i] = math.MaxInt
		low[i] = math.MaxInt
	}

	// Tarjan's algorithm
	criticals := [][]int{}
	var bridges [][]int
	for i := 0; i < n; i++ {
		if !visited[i] {
			bridges, time = dfsUtil(graph, i, time, parent, disc, low, visited)
			criticals = append(criticals, bridges...)
		}
	}

	return criticals
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
