package strobogrammaticnumber

func isStrobogrammatic(num string) bool {
	// use two pointer to check if chars match the

	ptr1 := 0
	ptr2 := len(num) - 1
	// while prt2 > prt1, do check
	for ptr2 >= ptr1 {
		// check if two chars match
		if !isMatch(num[ptr1], num[ptr2]) {
			return false
		}

		ptr1++
		ptr2--
	}
	return true
}

func isMatch(char1, char2 byte) bool {
	// 0 & 0
	// 1 & 1
	// 6 & 9
	// 8 & 8
	// 9 & 6
	if char1 == '0' && char2 == '0' {
		return true
	} else if char1 == '1' && char2 == '1' {
		return true
	} else if char1 == '6' && char2 == '9' {
		return true
	} else if char1 == '8' && char2 == '8' {
		return true
	} else if char1 == '9' && char2 == '6' {
		return true
	}
	return false
}
