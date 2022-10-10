package reconstructitinerary

import "sort"

type vertices []string

func (vs vertices) Len() int {
	return len(vs)
}

func (vs vertices) Less(i, j int) bool {
	return vs[i] < vs[j]
}

func (vs vertices) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}

func findItinerary(tickets [][]string) []string {
	// This is a Eulerian path problem.
	// We need to find a path that visits all edges exactly once starting from JFK.
	// if there are mulitple paths, we choose the path in smallest lexical order.

	// use Hierholzer’s Algorithm to solve this in O(V+E) time

	// build adjacency list
	adjLists := map[string][]string{}
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]

		if _, ok := adjLists[from]; !ok {
			adjLists[from] = []string{to}
		} else {
			adjLists[from] = append(adjLists[from], to)
		}
	}

	// each adjacency list is sorted in lexical order
	for from := range adjLists {
		sort.Sort(vertices(adjLists[from]))
	}

	// start from JFK
	path := []string{"JFK"}
	result := []string{}
	for len(path) > 0 {
		from := path[len(path)-1]
		if len(adjLists[from]) == 0 {
			// nowhere to go, add to result
			result = append(result, from)

			// start backtracking
			path = path[:len(path)-1]
		} else {
			// go to next vertex
			to := adjLists[from][0]
			adjLists[from] = adjLists[from][1:]

			path = append(path, to)
		}
	}

	// reverse result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
