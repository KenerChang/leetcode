package movingaveragefromdatastream

type MovingAverage struct {
	size  int
	kept  []int
	total int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (ma *MovingAverage) Next(val int) float64 {
	if len(ma.kept) >= ma.size {
		// pop first
		ma.total -= ma.kept[0]
		ma.total += val

		ma.kept = append(ma.kept[1:], val)
	} else {
		ma.total += val
		ma.kept = append(ma.kept, val)
	}

	return float64(ma.total) / float64(len(ma.kept))
}
