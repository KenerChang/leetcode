package racecar

import (
	"math"
)

func racecar(target int) int {
	// use dynamic programming to find minimum steps to reach target

	// init dp
	dp := make([]int, max(3, target+1))
	dp[0] = 0
	dp[1] = 1
	dp[2] = 4
	for i := 3; i <= target; i++ {
		dp[i] = math.MaxInt32
	}

	if target <= 2 {
		return dp[target]
	}

	for t := 3; t <= target; t++ {
		// if t is 2^x - 1, then it has minimum steps of x
		bitLength := getBitsCount(t)
		// fmt.Printf("t: %d, bit length: %d\n", t, bitLength)
		if t == (1<<bitLength)-1 {
			dp[t] = bitLength
			continue
		}

		// dp transition formula
		// min steps might come from 2 situations

		// first we reach bitLength -1 steps, then we turn back for some steps
		// then turn back again to reach t
		// in this case min step of t is dp[t-2^(k-1)+2^(j)] + k-1+j+2
		distanceOfBitLengthMinus1Steps := int(math.Pow(2, float64(bitLength-1))) - 1
		for j := 0; j < bitLength-1; j++ {
			distanceOfJSteps := int(math.Pow(2, float64(j))) - 1
			dp[t] = min(dp[t], dp[t-distanceOfBitLengthMinus1Steps+distanceOfJSteps]+bitLength-1+j+2)
		}

		// second, we go across t and turn back for some steps
		distanceOfBitLengthMinuse1Steps := int(math.Pow(2, float64(bitLength))) - 1
		if distanceOfBitLengthMinuse1Steps-t < t {
			dp[t] = min(dp[t], dp[distanceOfBitLengthMinuse1Steps-t]+bitLength+1)
		}
	}

	return dp[target]
}

func getBitsCount(n int) int {
	nFloat64 := float64(n)
	power := math.Log2(nFloat64)

	if nFloat64 == math.Pow(2, power) {
		return int(power) + 1
	}
	return int(math.Ceil(power))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
