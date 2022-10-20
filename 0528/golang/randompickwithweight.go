package randompickwithweight

import (
	"math/rand"
	"time"
)

type Solution struct {
	weights []float64
}

func Constructor(w []int) Solution {
	var sum int

	for _, weight := range w {
		sum += weight
	}

	weights := make([]float64, len(w))
	sumFloat64 := float64(sum)

	weights[0] = float64(w[0]) / sumFloat64
	for i := 1; i < len(w); i++ {
		weights[i] = float64(w[i])/sumFloat64 + weights[i-1]
	}

	rand.Seed(time.Now().Unix())
	return Solution{
		weights: weights,
	}
}

func (this *Solution) PickIndex() int {
	r := rand.Float64()

	return this.search(0, len(this.weights)-1, r)
}

func (this *Solution) search(l, r int, w float64) int {
	for l < r {
		middle := (l + r) / 2

		if middle == 0 {
			if this.weights[middle] >= w {
				return middle
			} else {
				return r
			}
		} else {
			if this.weights[middle] >= w && this.weights[middle-1] < w {
				return middle
			} else if w > this.weights[middle] {
				// search right
				l = middle + 1
			} else {
				// search left
				r = middle - 1
			}
		}
	}
	return l
}
