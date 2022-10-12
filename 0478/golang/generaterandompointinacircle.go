package generaterandompointinacircle

import (
	"math/rand"
	"time"
)

type Solution struct {
	r float64
	x float64
	y float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	rand.Seed(time.Now().Unix())
	return Solution{radius, x_center, y_center}
}

func (this *Solution) RandPoint() []float64 {
	x0 := this.x - this.r
	y0 := this.y - this.r
	for {
		x := x0 + rand.Float64()*this.r*2
		y := y0 + rand.Float64()*this.r*2
		if (x-this.x)*(x-this.x)+(y-this.y)*(y-this.y) <= this.r*this.r {
			return []float64{x, y}
		}
	}
}
