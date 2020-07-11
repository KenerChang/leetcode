package numberof1bits

import (
	"math"
)

func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}

	if num == math.MaxUint32 {
		return 32
	}

	var count int
	for num > 0 {
		count += int(num % 2)
		num = num >> 1
	}
	return count
}
