package topkfrequentwords

import "sort"

type sortableWords []string

func (w sortableWords) Len() int {
	return len(w)
}

func (w sortableWords) Less(i, j int) bool {
	return w[i] < w[j]
}

func (w sortableWords) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func topKFrequent(words []string, k int) []string {
	// Since the words length is bounded, we can use bucket sort to get top k frequent

	wordCountMap := map[string]int{}
	countToWordsMap := map[int]map[string]bool{}

	// count words
	maxCount := 0
	for _, word := range words {
		originalCount := wordCountMap[word]
		if originalCount == 0 {
			wordCountMap[word] = 1

			if countToWordsMap[1] == nil {
				countToWordsMap[1] = map[string]bool{}
			}
			countToWordsMap[1][word] = true
		} else {
			wordCountMap[word] = originalCount + 1

			if countToWordsMap[originalCount+1] == nil {
				countToWordsMap[originalCount+1] = map[string]bool{}
			}
			countToWordsMap[originalCount+1][word] = true

			delete(countToWordsMap[originalCount], word)
		}
		maxCount = max(maxCount, originalCount+1)
	}

	//
	result := []string{}
	remain := k
	i := maxCount
	for remain > 0 {
		if countToWordsMap[i] != nil {
			ws := []string{}
			for word := range countToWordsMap[i] {
				ws = append(ws, word)
			}

			sort.Sort(sortableWords(ws))
			if remain >= len(ws) {
				result = append(result, ws...)
				remain -= len(ws)
			} else {
				result = append(result, ws[:remain]...)
				remain = 0
			}
		}
		i--
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
