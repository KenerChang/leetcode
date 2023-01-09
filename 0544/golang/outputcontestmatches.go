package outputcontestmatches

import (
	"fmt"
	"strconv"
)

func findContestMatch(n int) string {
	q := []string{}
	for i := 1; i <= n; i++ {
		q = append(q, strconv.Itoa(i))
	}

	return recursive(q)
}

func recursive(q []string) string {
	if len(q) == 1 {
		return q[0]
	}

	nextRound := []string{}
	for i := 0; i < len(q)/2; i++ {
		match := fmt.Sprintf("(%s,%s)", q[i], q[len(q)-i-1])
		nextRound = append(nextRound, match)
	}

	return recursive(nextRound)
}
