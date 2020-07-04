package largestnumber

import (
	// "fmt"
	"sort"
	"strconv"
)

type SortableStrings []string

func (ss SortableStrings) Len() int {
	return len(ss)
}

func (ss SortableStrings) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss SortableStrings) Less(i, j int) bool {
	return ss[i]+ss[j] < ss[j]+ss[i]
}

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	buckets := make([]SortableStrings, 10)

	var fd int
	for _, num := range nums {
		fd = firstDigit(num)
		buckets[fd] = append(buckets[fd], strconv.Itoa(num))
	}

	// sort each bucket
	for _, bucket := range buckets {
		sort.Sort(bucket)
	}

	var bucket SortableStrings
	var result string
	var allZero bool
	for i := 9; i >= 0; i-- {
		bucket = buckets[i]
		if i > 0 {
			allZero = allZero || len(bucket) > 0
		}

		for j := len(bucket) - 1; j >= 0; j-- {
			result += bucket[j]
		}
	}

	if !allZero {
		result = "0"
	}
	return result
}

func firstDigit(num int) int {
	var digit int = num
	for digit >= 10 {
		digit = digit / 10
	}
	return digit
}
