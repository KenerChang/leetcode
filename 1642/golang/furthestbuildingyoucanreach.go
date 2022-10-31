package furthestbuildingyoucanreach

import (
	"sort"
)

type byHeight [][]int

func (b byHeight) Len() int {
	return len(b)
}

func (b byHeight) Less(i, j int) bool {
	return b[i][0] < b[j][0]
}

func (b byHeight) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func canReach(climbs [][]int, target, bricks, ladders int) bool {
	for _, climb := range climbs {
		if climb[1] > target {
			continue
		}

		if bricks >= climb[0] {
			bricks -= climb[0]
		} else if ladders > 0 {
			ladders--
		} else {
			return false
		}
	}

	return true
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	// We can solve this with binary search
	// We can use binary search becasuse we can view the buildings as two group: can reach & can not reach
	// the can reach group comes before the can not reach group
	// so we can use binary search to find the first building in can not reach group

	// collect climbs
	climbs := [][]int{}
	for i := 0; i < len(heights)-1; i++ {
		if heights[i] >= heights[i+1] {
			continue
		}

		climbs = append(climbs, []int{heights[i+1] - heights[i], i + 1})
	}

	// sort climbs
	sort.Sort(byHeight(climbs))

	l, r := 0, len(heights)
	for l < r {
		mid := (l + r) / 2
		if canReach(climbs, mid, bricks, ladders) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l - 1
}
