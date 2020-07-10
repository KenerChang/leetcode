package reversebits

import (
	// "fmt"
	"math"
)

func reverseBits(num uint32) uint32 {
	if num == 0 || num == math.MaxUint32 {
		return num
	}

	remain := num
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		result = result << 1
		result = result + remain%2
		remain = remain >> 1
	}
	return result
}
