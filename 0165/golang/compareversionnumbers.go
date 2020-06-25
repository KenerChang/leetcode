package compareversionnumbers

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	if version1 == version2 {
		return 0
	}

	snippets1, snippets2 := strings.Split(version1, "."), strings.Split(version2, ".")
	maxLen := max(len(snippets1), len(snippets2))

	var v1, v2 int
	for i := 0; i < maxLen; i++ {
		if i >= len(snippets1) {
			v1 = 0
		} else {
			v1, _ = strconv.Atoi(snippets1[i])
		}

		if i >= len(snippets2) {
			v2 = 0
		} else {
			v2, _ = strconv.Atoi(snippets2[i])
		}

		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	return 0
}

func max(num1, num2 int) int {
	if num1 >= num2 {
		return num1
	}
	return num2
}
