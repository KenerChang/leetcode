package theskylineproblem

import (
	"container/heap"
	"math"
	"sort"
)

type Positions []Position

func (ps Positions) Len() int {
	return len(ps)
}

func (ps Positions) Less(i, j int) bool {
	return ps[i].X < ps[j].X
}

func (ps Positions) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type PositionPriorityQueue []Position

func (pq PositionPriorityQueue) Len() int { return len(pq) }

func (pq PositionPriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return math.Abs(float64(pq[i].Height)) > math.Abs(float64(pq[j].Height))
}

func (pq PositionPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PositionPriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Position))
}

func (pq *PositionPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

type Position struct {
	X      int
	Height int
}

func getSkyline(buildings [][]int) [][]int {
	// use two priority queue + Sweep Line
	// first get each position

	positions := []Position{}
	for _, building := range buildings {
		positions = append(positions, Position{
			X:      building[0],
			Height: building[2],
		}, Position{
			X:      building[1],
			Height: -building[2],
		})
	}

	// sort positions
	sort.Sort(Positions(positions))

	// use two pritority to keep live & past positions
	live := PositionPriorityQueue{}
	heap.Init(&live)
	past := PositionPriorityQueue{}
	heap.Init(&past)

	skylines := [][]int{}
	i := 0
	for i < len(positions) {
		position := positions[i]

		x := position.X

		// handle all position at x
		for i < len(positions) && positions[i].X == x {
			p := positions[i]
			if p.Height > 0 {
				// a start point
				heap.Push(&live, p)
			} else {
				// it is a end position, check if we need to remove past positions
				heap.Push(&past, p)
			}
			i++
		}

		for past.Len() > 0 && past[0].Height == -live[0].Height {
			heap.Pop(&past)
			heap.Pop(&live)
		}

		var maxHeight int
		if live.Len() > 0 {
			maxHeight = live[0].Height
		}

		// if maxHeight changed
		if len(skylines) == 0 || skylines[len(skylines)-1][1] != maxHeight {
			skylines = append(skylines, []int{x, maxHeight})
		}
	}

	return skylines
}
