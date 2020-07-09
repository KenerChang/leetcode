package rotatearray

// import (
// 	"fmt"
// )

func rotate(nums []int, k int) {
	nLen := len(nums)
	realK := k % nLen
	if realK == 0 {
		return
	}

	start := nLen - realK
	target := append(nums[start:nLen], nums[0:start]...)

	for i := 0; i < nLen; i++ {
		nums[i] = target[i]
	}
}
