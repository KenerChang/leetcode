package analyzeuserwebsitevisitpattern

import (
	"container/heap"
	"fmt"
	"sort"
)

var (
	_ sort.Interface = visitHistories{}
	_ heap.Interface = &visitedPatterns{}
)

type visitedPattern struct {
	Site1 string
	Site2 string
	Site3 string
	Count int
}

type visitedPatterns []visitedPattern

func (ps visitedPatterns) Len() int {
	return len(ps)
}

func (ps visitedPatterns) Less(i, j int) bool {
	if ps[i].Count != ps[j].Count {
		return ps[i].Count > ps[j].Count
	}

	if ps[i].Site1 != ps[j].Site1 {
		return ps[i].Site1 < ps[j].Site1
	}

	if ps[i].Site2 != ps[j].Site2 {
		return ps[i].Site2 < ps[j].Site2
	}

	if ps[i].Site3 != ps[j].Site3 {
		return ps[i].Site3 < ps[j].Site3
	}

	return true
}

func (ps visitedPatterns) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps *visitedPatterns) Pop() interface{} {
	old := *ps
	size := len(old)

	head := old[size-1]
	*ps = old[0 : size-1]

	return head
}

func (ps *visitedPatterns) Push(x interface{}) {
	*ps = append(*ps, x.(visitedPattern))
}

type visitHistory struct {
	Name      string
	Timestamp int
	Website   string
}

type visitHistories []visitHistory

func (h visitHistories) Len() int {
	return len(h)
}

func (h visitHistories) Less(i, j int) bool {
	return h[i].Timestamp < h[j].Timestamp
}

func (h visitHistories) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func generatePatterns(websites []string) [][]string {
	visited := map[string]bool{}

	results := [][]string{}
	for i := 0; i < len(websites); i++ {
		for j := i + 1; j < len(websites); j++ {
			for k := j + 1; k < len(websites); k++ {
				key := fmt.Sprintf("%s_%s_%s", websites[i], websites[j], websites[k])

				if _, ok := visited[key]; !ok {
					visited[key] = true

					results = append(results, []string{websites[i], websites[j], websites[k]})
				}
			}
		}
	}
	return results
}

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	// sort visited websites by timestamp
	histories := []visitHistory{}
	for i := 0; i < len(username); i++ {
		histories = append(histories, visitHistory{
			Name:      username[i],
			Timestamp: timestamp[i],
			Website:   website[i],
		})
	}
	sort.Sort(visitHistories(histories))

	// collect users visited websites
	historiesOfUsers := map[string][]string{}
	for _, history := range histories {
		if _, ok := historiesOfUsers[history.Name]; !ok {
			historiesOfUsers[history.Name] = []string{}
		}

		historiesOfUsers[history.Name] = append(historiesOfUsers[history.Name], history.Website)
	}

	// generate patterns for each user
	patternsOfUsers := map[string][][]string{}
	for user := range historiesOfUsers {
		patternsOfUsers[user] = generatePatterns(historiesOfUsers[user])
	}

	// reduce patterns
	patterns := map[string]visitedPattern{}
	for user := range patternsOfUsers {
		for _, p := range patternsOfUsers[user] {
			key := fmt.Sprintf("%s_%s_%s", p[0], p[1], p[2])

			pattern, ok := patterns[key]
			if ok {
				pattern.Count++
				patterns[key] = pattern
			} else {
				patterns[key] = visitedPattern{
					Site1: p[0],
					Site2: p[1],
					Site3: p[2],
					Count: 1,
				}
			}
		}
	}

	// sort patterns (use max heap)
	h := &visitedPatterns{}
	heap.Init(h)

	for _, pattern := range patterns {
		heap.Push(h, pattern)
	}

	result := heap.Pop(h).(visitedPattern)

	return []string{result.Site1, result.Site2, result.Site3}
}
