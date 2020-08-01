package shortestpalindrome

func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	bytes := []byte(s)
	var left, right int
	var reset bool = true
	for left >= 0 || right < len(bytes) {
		if reset {
			even := (len(bytes) % 2) == 0
			middle := len(bytes) / 2
			if even {
				right = middle
				left = right - 1
			} else {
				right = middle + 1
				left = middle - 1
			}
			reset = false
		} else {
			if bytes[left] == '#' {
				bytes[left] = bytes[right]
				left--
				right++
			} else if bytes[left] == bytes[right] {
				left--
				right++
			} else {
				reset = true
				bytes = append([]byte{'#'}, bytes...)
			}
		}
	}

	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i] = bytes[j]
		i++
		j--
	}
	return string(bytes)
}
