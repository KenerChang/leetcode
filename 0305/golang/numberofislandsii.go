package numberofislandsii

import "fmt"

type DisjointSet struct {
	X            int
	Y            int
	Representive *DisjointSet
	Rank         int
}

func (s *DisjointSet) Find() *DisjointSet {
	// return reresentive of the disjointSet
	if s.Representive == s {
		return s.Representive
	}

	s.Representive = s.Representive.Find()
	return s.Representive
}

func (s *DisjointSet) Union(x *DisjointSet) {
	if s.Rank < x.Rank {
		s.Representive = x
	} else if s.Rank > x.Rank {
		x.Representive = s
	} else {
		// if s.Rank equals x.Rank, set s's representive to x
		// and increase x.Rank
		s.Representive = x
		x.Rank++
	}
}

func disjointSetFromPosition(position []int) *DisjointSet {
	set := &DisjointSet{
		X:    position[0],
		Y:    position[1],
		Rank: 1,
	}

	set.Representive = set
	return set
}

func numIslands2(m int, n int, positions [][]int) []int {
	// iterate all position
	visited := map[string]*DisjointSet{}
	results := []int{}
	setCount := 0
	for _, position := range positions {
		x, y := position[0], position[1]
		pKey := fmt.Sprintf("%d,%d", x, y)

		if _, ok := visited[pKey]; ok {
			results = append(results, setCount)
			continue
		}

		setP := disjointSetFromPosition(position)
		mergeCount := 0

		// check if upper grid is a island
		upperGridKey := fmt.Sprintf("%d,%d", x-1, y)
		if s, ok := visited[upperGridKey]; ok && s.Find() != setP.Find() {
			// fmt.Printf("upper: s.Find: %d,%d, setP.Find: %d,%d\n", s.Representive.X, s.Representive.Y, setP.Representive.X, setP.Representive.Y)
			s.Representive.Union(setP.Representive)
			mergeCount++
		}

		// check if left grid is an island
		leftGridKey := fmt.Sprintf("%d,%d", x, y-1)
		if s, ok := visited[leftGridKey]; ok && s.Find() != setP.Find() {
			// fmt.Printf("left: s.Find: %d,%d, setP.Find: %d,%d\n", s.Representive.X, s.Representive.Y, setP.Representive.X, setP.Representive.Y)
			s.Representive.Union(setP.Representive)
			mergeCount++
		}

		// check if bottom grid is an island
		bottomGridKey := fmt.Sprintf("%d,%d", x+1, y)
		if s, ok := visited[bottomGridKey]; ok && s.Find() != setP.Find() {
			// fmt.Printf("bottom: s.Find: %d,%d, setP.Find: %d,%d\n", s.Representive.X, s.Representive.Y, setP.Representive.X, setP.Representive.Y)
			s.Representive.Union(setP.Representive)
			mergeCount++
		}

		// check if right grid is an island
		rightGridKey := fmt.Sprintf("%d,%d", x, y+1)
		if s, ok := visited[rightGridKey]; ok && s.Find() != setP.Find() {
			// fmt.Printf("right: s.Find: %d,%d, setP.Find: %d,%d\n", s.Representive.X, s.Representive.Y, setP.Representive.X, setP.Representive.Y)
			s.Representive.Union(setP.Representive)
			mergeCount++
		}

		setCount = setCount - mergeCount + 1
		results = append(results, setCount)

		visited[fmt.Sprintf("%d,%d", x, y)] = setP
	}

	return results
}
