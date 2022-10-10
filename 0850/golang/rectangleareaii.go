package rectangleareaii

import (
	"sort"
)

const (
	m          = 1e9 + 7
	typeadd    = -1
	typeremove = 1
)

type position struct {
	x     int
	ys    int // start position in y direction
	ye    int // end position in y direction
	ptype int // start or end
	id    int
}

type positions []position

func (ps positions) Len() int {
	return len(ps)
}

func (ps positions) Less(i, j int) bool {
	if ps[i].x == ps[j].x {
		return ps[i].ptype < ps[j].ptype
	}
	return ps[i].x < ps[j].x
}

func (ps positions) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type ysegment struct {
	start int
	end   int
}

type ysegments []ysegment

func (ys ysegments) Len() int {
	return len(ys)
}

func (ys ysegments) Less(i, j int) bool {
	return ys[i].start < ys[j].start
}

func (ys ysegments) Swap(i, j int) {
	ys[i], ys[j] = ys[j], ys[i]
}

func (ys ysegments) Merged() [][]int {
	merged := [][]int{}

	if len(ys) == 0 {
		return merged
	}

	current := []int{ys[0].start, ys[0].end}
	for _, s := range ys {
		if s.start > current[1] {
			merged = append(merged, current)

			current = []int{s.start, s.end}
		} else {
			current[1] = max(current[1], s.end)
		}
	}

	merged = append(merged, current)
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mod(a, b int) int {
	return (a + b) % m
}

func rectangleArea(rs [][]int) int {
	// This is combination of sweep line & merge interval
	// we can sweep in x direction, and get how many y in current x and last x

	// get all positions
	// and sort rectangles by its start position and end position in x direction
	ps := make(positions, len(rs)*2)
	for i, r := range rs {
		ps[2*i] = position{r[0], r[1], r[3], typeadd, i}
		ps[2*i+1] = position{r[2], r[1], r[3], typeremove, i}
	}
	sort.Sort(ps)

	// start sweep
	lastX := 0
	ysegs := []ysegment{}
	var result int
	for _, p := range ps {
		sort.Sort(ysegments(ysegs))
		merged := ysegments(ysegs).Merged()

		width := p.x - lastX
		for _, m := range merged {
			result = mod(result, width*(m[1]-m[0]))
		}

		if p.ptype == typeadd {
			ysegs = append(ysegs, ysegment{p.ys, p.ye})
		} else {
			// remove from set
			for i, s := range ysegs {
				if s.start == p.ys && s.end == p.ye {
					ysegs = append(ysegs[:i], ysegs[i+1:]...)
					break
				}
			}
		}

		lastX = p.x
	}

	return result
}
