package longestincreasingpathinamatrix

import "fmt"

func longestIncreasingPath(matrix [][]int) int {
	// we can view each cell in the matrix as a nodie in a graph
	// and we define an edge between two nodes as a -> b if position a < position b
	// then this is a DAG and we can solve path problem by topological sort

	if len(matrix) == 0 {
		return 0
	}

	if len(matrix) == 1 && len(matrix[0]) == 1 {
		return 1
	}

	// extract relationship between nodes
	relations := extract(matrix)
	// fmt.Printf("relations: %v\n", relations)

	// build graph
	indegrees, outdegrees := buildGraph(relations)
	// fmt.Printf("indegrees: %v, outdegrees: %v\n", indegrees, outdegrees)

	// topological sort (DFS)
	var longest int
	cache := map[string]int{}
	for node := range indegrees {
		if indegrees[node] == 0 {
			// fmt.Printf("node: %+v\n", node)
			longest = max(longest, dfs(node, indegrees, outdegrees, cache))
		}
	}

	return longest
}

func buildGraph(relations [][]string) (map[string]int, map[string][]string) {
	indegrees := map[string]int{}
	outdegrees := map[string][]string{}
	for _, relation := range relations {
		start := relation[0]
		end := relation[1]

		if _, ok := indegrees[start]; !ok {
			indegrees[start] = 0
		}
		indegrees[end]++

		if _, ok := outdegrees[start]; !ok {
			outdegrees[start] = []string{end}
		} else {
			outdegrees[start] = append(outdegrees[start], end)
		}
	}
	return indegrees, outdegrees
}

func dfs(node string, indegrees map[string]int, outdegrees map[string][]string, cache map[string]int) int {
	if val, ok := cache[node]; ok {
		return val
	}

	nextNodes := outdegrees[node]

	pathLength := 1
	for _, nextNode := range nextNodes {
		pathLength = max(pathLength, 1+dfs(nextNode, indegrees, outdegrees, cache))
	}

	cache[node] = pathLength
	return pathLength
}

func extract(matrix [][]int) [][]string {
	relations := [][]string{}
	for i, row := range matrix {
		for j, val := range row {
			// check up
			if i > 0 && matrix[i-1][j] > val {
				relation := []string{
					fmt.Sprintf("%d,%d", i, j),
					fmt.Sprintf("%d,%d", i-1, j),
				}
				relations = append(relations, relation)
			}

			// check down
			if i < len(matrix)-1 && matrix[i+1][j] > val {
				relation := []string{
					fmt.Sprintf("%d,%d", i, j),
					fmt.Sprintf("%d,%d", i+1, j),
				}
				relations = append(relations, relation)
			}

			// check left
			if j > 0 && matrix[i][j-1] > val {
				relation := []string{
					fmt.Sprintf("%d,%d", i, j),
					fmt.Sprintf("%d,%d", i, j-1),
				}
				relations = append(relations, relation)
			}

			// check right
			if j < len(row)-1 && matrix[i][j+1] > val {
				relation := []string{
					fmt.Sprintf("%d,%d", i, j),
					fmt.Sprintf("%d,%d", i, j+1),
				}
				relations = append(relations, relation)
			}
		}
	}

	return relations
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
