package numberofconnectedcomponentsinanundirectedgraph

func countComponents(n int, edges [][]int) int {
	// use dfs to solve this problem

	// build adj
	neighbors := make([][]int, n)
	for _, edge := range edges {
		if neighbors[edge[0]] == nil {
			neighbors[edge[0]] = []int{}
		}

		if neighbors[edge[1]] == nil {
			neighbors[edge[1]] = []int{}
		}

		neighbors[edge[0]] = append(neighbors[edge[0]], edge[1])
		neighbors[edge[1]] = append(neighbors[edge[1]], edge[0])
	}

	// do dfs marked vertices as marked
	visited := make([]bool, n)
	var componentCount int
	for i := 0; i < n; i++ {
		if !visited[i] {
			componentCount++
			dfs(neighbors, visited, i)
		}
	}

	return componentCount
}

func dfs(neighbors [][]int, visited []bool, i int) {
	ineighbors := neighbors[i]

	visited[i] = true
	for _, neighbor := range ineighbors {
		if !visited[neighbor] {
			dfs(neighbors, visited, neighbor)
		}
	}
}
