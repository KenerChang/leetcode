package kclosestpointstoorigin

type Point []int

func (p Point) distance() int {
	return p[0]*p[0] + p[1]*p[1]
}

func kClosest(points [][]int, k int) [][]int {
	// since we are finding the top k closest points, we can use quick select
	if len(points) == k {
		return points
	}

	quickSelect(points, 0, len(points)-1, k)

	return points[:k]
}

func quickSelect(ps [][]int, l, r, k int) {
	if l-r == 1 {
		return
	}

	// choose last point as pivot
	pivotDistance := Point(ps[r]).distance()
	i := l
	for j := l; j < r; j++ {
		if Point(ps[j]).distance() <= pivotDistance {
			ps[i], ps[j] = ps[j], ps[i]
			i++
		}
	}

	// swap pivot
	ps[i], ps[r] = ps[r], ps[i]

	if i-l == k-1 {
		return
	}

	if i-l > k-1 {
		// search left
		quickSelect(ps, l, i-1, k)
	} else {
		// search right
		quickSelect(ps, i+1, r, k-i+l-1)
	}
}
