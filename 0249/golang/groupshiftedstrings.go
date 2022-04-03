package groupshiftedstrings

import "fmt"

func groupStrings(strings []string) [][]string {
	groups := map[string][]string{}

	// for each string in Strings
	for _, s := range strings {
		// hash string to key
		key := hashString(s)

		// if key does not exist, create a new string group
		if _, ok := groups[key]; !ok {
			groups[key] = []string{}
		}

		// put the string to group
		groups[key] = append(groups[key], s)
	}

	results := [][]string{}
	for _, group := range groups {
		results = append(results, group)
	}

	return results
}

func hashString(s string) string {
	if len(s) == 1 {
		return "_"
	}

	key := ""
	for i := 1; i < len(s); i++ {
		diff := int(s[i]) - int(s[i-1])
		if diff < 0 {
			diff += 26
		}

		key = fmt.Sprintf("%s_%d", key, diff)
	}

	return key
}
