package shortestworddistanceii

import (
	"fmt"
	"math"
)

type WordDistance struct {
	wordsDict      []string
	wordsPositions map[string][]int
	cache          map[string]int
}

func Constructor(wordsDict []string) WordDistance {
	wd := WordDistance{
		wordsDict:      wordsDict,
		cache:          map[string]int{},
		wordsPositions: map[string][]int{},
	}

	for i, word := range wordsDict {
		if _, ok := wd.wordsPositions[word]; !ok {
			wd.wordsPositions[word] = []int{}
		}

		wd.wordsPositions[word] = append(wd.wordsPositions[word], i)
	}

	return wd
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	// check if counted
	key := fmt.Sprintf("%s_%s", word1, word2)
	if val, ok := this.cache[key]; ok {
		return val
	}

	// if not counted, count shourtest distance
	distance := this.shortestDistance(word1, word2)

	// save to cache
	key1 := fmt.Sprintf("%s_%s", word1, word2)
	key2 := fmt.Sprintf("%s_%s", word2, word1)

	this.cache[key1] = distance
	this.cache[key2] = distance

	return distance
}

func (this *WordDistance) shortestDistance(word1, word2 string) int {
	// iterate word1 & word2 position list

	ptr1, ptr2 := 0, 0
	list1 := this.wordsPositions[word1]
	list2 := this.wordsPositions[word2]
	distance := math.MaxInt32
	keep := true
	for keep {
		// while there are remained element in word1 and word2 list
		// count shortest distance
		num1 := list1[ptr1]
		num2 := list2[ptr2]

		if num1 > num2 {
			distance = min(distance, num1-num2)
			ptr2++
		}

		if num1 < num2 {
			distance = min(distance, num2-num1)
			ptr1++
		}

		// touch end
		if ptr1 >= len(list1) && ptr2 >= len(list2) {
			keep = false
		}

		// no need to keep iterating list2
		if ptr1 >= len(list1) && num1 < num2 {
			keep = false
		}

		// no need to keep iterating list 1
		if ptr2 >= len(list2) && num2 < num1 {
			keep = false
		}
	}

	return distance
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */
