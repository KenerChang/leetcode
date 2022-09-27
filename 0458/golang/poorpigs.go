package poorpigs

import (
	"math"
)

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	// This is a qbit problem, so all we need to find is that how many states a pig can have
	// States of a pig can have depends on how many times we can test it
	// for example, if minutesToDie / minutesToTest = 1, then the pig can only have 2 states: alive or dead
	// and if minutesToTest / minutesToDie = 2, then the pig can have 3 states: alive, dead at first round, dead at second round
	// for above two examples, we can know that states = minutesToTest / minutesToDie + 1
	// and all we need to do is to find the minimum number of s (pigs) that states ^ s >= buckets

	states := minutesToTest/minutesToDie + 1
	pigs := math.Log2(float64(buckets)) / math.Log2(float64(states))

	return int(math.Ceil(pigs))
}
