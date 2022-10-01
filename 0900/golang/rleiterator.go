package rleiterator

type RLEIterator struct {
	currPos   int
	currNum   int
	currCount int
	encoding  []int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{
		currPos:   0,
		currNum:   encoding[1],
		currCount: encoding[0],
		encoding:  encoding,
	}
}

func (iter *RLEIterator) Next(n int) int {
	if iter.currPos >= len(iter.encoding) {
		return -1
	}

	if iter.currCount >= n {
		iter.currCount -= n
		return iter.currNum
	} else {
		remain := n
		for remain > iter.currCount {
			remain -= iter.currCount

			if remain == 0 {
				iter.currCount = 0
				break
			}

			iter.currPos += 2

			if iter.currPos >= len(iter.encoding) {
				break
			}

			iter.currCount = iter.encoding[iter.currPos]
			iter.currNum = iter.encoding[iter.currPos+1]
		}

		if iter.currPos >= len(iter.encoding) {
			return -1
		} else {
			iter.currCount -= remain
			return iter.currNum
		}
	}
}
