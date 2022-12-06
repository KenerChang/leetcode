package minimumheighttrees

func findMinHeightTrees(n int, edges [][]int) []int {
	// We can solve this problem by toplogical sort
	// we remove leaves in each iteration, the last nodes are the answer

	// edge case
	if n < 2 {
		return []int{0}
	}

	// build adj
	neighbors := map[int]map[int]bool{}
	for _, edge := range edges {
		if neighbors[edge[0]] == nil {
			neighbors[edge[0]] = map[int]bool{}
		}

		if neighbors[edge[1]] == nil {
			neighbors[edge[1]] = map[int]bool{}
		}

		neighbors[edge[0]][edge[1]] = true
		neighbors[edge[1]][edge[0]] = true
	}

	// get leaves
	leaves := []int{}
	for k, v := range neighbors {
		if len(v) == 1 {
			leaves = append(leaves, k)
		}
	}

	// remove leaves
	remainNodes := n
	for remainNodes > 2 {
		remainNodes -= len(leaves)

		newLeaves := []int{}
		for _, v := range leaves {
			vneighbors := neighbors[v]

			for neighbor := range vneighbors {
				delete(neighbors[neighbor], v)

				if len(neighbors[neighbor]) == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}

		leaves = newLeaves
	}

	return leaves
}
