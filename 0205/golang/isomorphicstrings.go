package isomorphicstrings

import (
// "fmt"
)

func isIsomorphic(s string, t string) bool {
	if s == t {
		return true
	}
	sMap := map[rune]rune{}
	tMap := map[rune]rune{}

	for i, chars := range s {
		nowChart := rune(t[i])

		smapt, founds := sMap[chars]
		tmaps, foundt := tMap[nowChart]
		if !founds && !foundt {
			// pair does not exist
			sMap[chars] = nowChart
			tMap[nowChart] = chars
			continue
		}

		if smapt != nowChart || tmaps != chars {
			return false
		}
		// fmt.Printf("nowChart: %c\n", nowChart)
	}
	return true
}
