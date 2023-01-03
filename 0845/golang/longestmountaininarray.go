package longestmountaininarray

const (
	notBegin = iota
	ascending
	descending
)

func longestMountain(arr []int) int {
	if len(arr) < 3 {
		return 0
	}

	state := notBegin
	var longest int
	var begin int
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			if state == descending {
				// might be end of a mountain and start to be a new mountain
				longest = max(longest, i-begin)
				begin = i - 1
				state = ascending
			} else if state == notBegin {
				begin = i - 1
				state = ascending
			}
		} else if arr[i] < arr[i-1] {
			if state == ascending {
				state = descending
			}
		} else {
			// arr[i] == arr[i-1]
			if state == descending {
				// end of a mountain
				longest = max(longest, i-begin)
				state = notBegin
			} else if state == ascending {
				state = notBegin
			}
		}
	}

	if state == descending {
		longest = max(longest, len(arr)-begin)
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
