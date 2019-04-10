package sqrt

// mySqrt return integer part of squre root
// time complexity of mySqrt is O(logN)
// we do bit shift to approximate the square root
// if the squre root is larger than real square root
// we reset i and do again until square root is approximated
func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}

	result := 0
	i := 1
	for result*result <= x {
		current := result + i
		square := current * current
		if square == x {
			result += i
			break
		} else if square > x {
			if i == 1 {
				break
			}
			result += i >> 1
			i = 1
		} else {
			i = i << 1
		}
	}
	return result
}
