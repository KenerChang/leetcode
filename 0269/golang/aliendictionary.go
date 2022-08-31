package aliendictionary

import "fmt"

func alienOrder(words []string) string {
	// We can view alien order as a topological sort problem
	// so we need to build the DAG (directed acyclic graph)
	// before appling topological sort (either DFS / Kahan's).

	// if len(words) == 1 {
	// 	return words[0]
	// }

	// extract relations
	relations := [][]byte{}

	// extract relations from first characters
	relations = append(relations, extractFirstCharRelation(words))

	// extract relations from comparing two adjacent words
	wordRelations, err := extractWordRelations(words)
	if err != nil {
		return ""
	}
	relations = append(relations, wordRelations...)

	// build graphs
	// use two map to keep indegree & outdegree information
	indegrees, outdegrees := buildGraph(relations)

	// fill up chars not in relations
	for _, word := range words {
		for j := range word {
			char := word[j]
			if _, ok := indegrees[char]; !ok {
				indegrees[char] = 0
				outdegrees[char] = []byte{}
			}
		}
	}

	// do topological sort
	// use Kahan's algorithm here, every time indegree map is updated,
	// we add nodes n which indegree(n) is 0 to the queue
	queue := getZeroIndegreeChars(indegrees)
	var result string
	for len(queue) > 0 {
		char := queue[0]
		queue = queue[1:]

		result += string(char)
		for _, m := range outdegrees[char] {
			indegrees[m]--
		}

		queue = append(queue, getZeroIndegreeChars(indegrees)...)
	}

	// check if cycle existed
	if len(indegrees) > 0 {
		return ""
	}

	return result
}

func extractFirstCharRelation(words []string) []byte {
	firstCharRelation := []byte{}
	// charExisted := map[byte]bool{}
	for _, word := range words {
		firstChar := word[0]
		// if _, ok := charExisted[firstChar]; !ok {

		// 	// charExisted[firstChar] = true
		// }
		firstCharRelation = append(firstCharRelation, firstChar)
	}

	return firstCharRelation
}

func extractWordRelations(words []string) ([][]byte, error) {
	existedRelations := map[string]bool{}
	relations := [][]byte{}
	for i := 0; i < len(words)-1; i++ {
		word, nextWord := words[i], words[i+1]

		firstDiffWord1, firstDiffWord2, err := findFirstDiffChar(word, nextWord)
		if err != nil {
			return relations, fmt.Errorf("invalid relation")
		}
		key := fmt.Sprintf("%s%s\n", string(firstDiffWord1), string(firstDiffWord2))

		if _, ok := existedRelations[key]; !ok {
			existedRelations[key] = true
			relations = append(relations, []byte{firstDiffWord1, firstDiffWord2})
		}
	}

	return relations, nil
}

func findFirstDiffChar(word1, word2 string) (byte, byte, error) {
	minWordLen := min(len(word1), len(word2))
	for i := 0; i < minWordLen; i++ {
		if word1[i] != word2[i] {
			return word1[i], word2[i], nil
		}
	}

	if len(word1) > len(word2) {
		return byte(' '), byte(' '), fmt.Errorf("invalid input")
	}

	return word1[minWordLen-1], word2[minWordLen-1], nil
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func buildGraph(relations [][]byte) (map[byte]int, map[byte][]byte) {
	indegree := map[byte]int{}
	outdegree := map[byte][]byte{}

	usedRelations := map[string]bool{}
	for _, relation := range relations {
		for i := 0; i < len(relation)-1; i++ {
			key := fmt.Sprintf("%s%s", string(relation[i]), string(relation[i+1]))

			if _, ok := usedRelations[key]; !ok {
				usedRelations[key] = true
				if relation[i] != relation[i+1] {
					indegree[relation[i+1]]++
					outdegree[relation[i]] = append(outdegree[relation[i]], relation[i+1])
				}

				if _, iok := indegree[relation[i]]; !iok {
					indegree[relation[i]] = 0
				}
			}
		}
	}

	return indegree, outdegree
}

func getZeroIndegreeChars(indegrees map[byte]int) []byte {
	zeroIndegreeChars := []byte{}
	for char, indegree := range indegrees {
		if indegree == 0 {
			zeroIndegreeChars = append(zeroIndegreeChars, char)
			delete(indegrees, char)
		}
	}

	return zeroIndegreeChars
}
