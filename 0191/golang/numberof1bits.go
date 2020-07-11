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
	for i := 0; i < 32; i++ {
		count += int(num % 2)
		num = num >> 1
	}
	return count
}
