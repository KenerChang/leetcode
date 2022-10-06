package sortcharactersbyfrequency

import "sort"

type pair [][]int // 0: char 1: frequency

func (p pair) Len() int {
	return len(p)
}

func (p pair) Less(i, j int) bool {
	return p[i][1] > p[j][1]
}

func (p pair) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func getIndex(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a')
	} else if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 26
	} else if c >= '0' && c <= '9' {
		return int(c-'0') + 52
	}

	return -1
}

func getChar(c int) byte {
	if c >= 0 && c <= 25 {
		return byte(c + 'a')
	} else if c >= 26 && c <= 51 {
		return byte(c - 26 + 'A')
	} else if c >= 52 && c <= 61 {
		return byte(c - 52 + '0')
	}

	return 0
}

func frequencySort(s string) string {
	pairs := make([][]int, 62) // a - z, A - Z, 0 - 9
	for i := 0; i < 62; i++ {
		pairs[i] = []int{i, 0}
	}

	for _, c := range s {
		i := getIndex(c)

		pairs[i][1]++
	}

	sort.Sort(pair(pairs))

	result := []byte{}
	for _, p := range pairs {
		if len(p) == 0 {
			continue
		}

		char := getChar(p[0])
		for i := 0; i < p[1]; i++ {
			result = append(result, char)
		}
	}

	return string(result)
}
