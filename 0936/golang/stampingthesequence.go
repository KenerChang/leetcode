package stampingthesequence

import (
	"strings"
)

func generateCandidates(stamp string) []string {
	set := map[string]bool{}
	for i := 0; i < len(stamp); i++ {
		for j := 0; j < len(stamp)-i; j++ {
			candidate := strings.Repeat("#", i)
			candidate += stamp[i : len(stamp)-j]
			candidate += strings.Repeat("#", j)

			set[candidate] = true
		}
	}

	results := []string{}
	for candidate := range set {
		results = append(results, candidate)
	}
	return results
}

func getSharpString(num int) string {
	bytes := make([]byte, num)
	for i := 0; i < len(bytes); i++ {
		bytes[i] = '#'
	}
	return string(bytes)
}

func movesToStamp(stamp string, target string) []int {
	// We can solve this problem in reverse order, which means we can find a path
	// to transform to s, for example, if target is "abcde", then we want to find
	// a path to transform "abcde" into "?????", then we can reverse the path as the answer.

	// To find such a path, we can generate all possible match candidates of the stamp
	// for example, if stamp is "abc", then we have match candidates {"abc", "ab#", "a#c", "#bc", "##c", "a##"}
	// then do candiates replacement until target becomes s

	// generate match candidates
	candidates := generateCandidates(stamp)
	// fmt.Printf("candidates: %v\n", candidates)

	// then we create s
	s := getSharpString(len(target))
	transformedString := getSharpString((len(stamp)))

	// fmt.Printf("s: %s, transformedString: %s", s, transformedString)

	pathes := []int{}
	for target != s {
		transfromed := false
		for _, candidate := range candidates {

			pos := strings.Index(target, candidate)
			// fmt.Printf("pos: %d, candidate: %s\n", pos, candidate)
			if pos != -1 {
				pathes = append(pathes, pos)
				target = strings.Replace(target, candidate, transformedString, 1)
				transfromed = true

				break
			}
		}

		// fmt.Printf("pathes: %v, transofrmed: %t, target: %s, s: %s, target == s: %t\n", pathes, transfromed, target, s, target == s)

		if !transfromed {
			// can not transform target to s
			return []int{}
		}
	}

	// reverse pathes
	i, j := 0, len(pathes)-1
	for i < j {
		pathes[i], pathes[j] = pathes[j], pathes[i]
		i++
		j--
	}

	return pathes
}
