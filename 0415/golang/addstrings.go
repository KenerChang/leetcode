package addstrings

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	}

	if num2 == "0" {
		return num1
	}

	var addOne int
	results := []int{}
	i := len(num1) - 1
	j := len(num2) - 1
	for i >= 0 && j >= 0 {
		d1 := int(num1[i] - '0')
		d2 := int(num2[j] - '0')

		sum := d1 + d2 + addOne
		results = append(results, sum%10)
		if sum >= 10 {
			addOne = 1
		} else {
			addOne = 0
		}

		i--
		j--
	}

	for i >= 0 {
		sum := int(num1[i]-'0') + addOne
		results = append(results, sum%10)

		if sum >= 10 {
			addOne = 1
		} else {
			addOne = 0
		}
		i--
	}

	for j >= 0 {
		sum := int(num2[j]-'0') + addOne
		results = append(results, sum%10)
		if sum >= 10 {
			addOne = 1
		} else {
			addOne = 0
		}

		j--
	}

	if addOne > 0 {
		results = append(results, 1)
	}

	var result string
	for i := len(results) - 1; i >= 0; i-- {
		result += strconv.Itoa(results[i])
	}

	return result
}
