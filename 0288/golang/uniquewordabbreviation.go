package uniquewordabbreviation

import "fmt"

type ValidWordAbbr struct {
	abbrMap map[string]map[string]struct{}
}

func getAbbr(word string) string {
	if len(word) <= 2 {
		return word
	}

	return fmt.Sprintf("%b%d%b", word[0], len(word)-2, word[len(word)-1])
}

func Constructor(dictionary []string) ValidWordAbbr {
	validator := ValidWordAbbr{
		abbrMap: map[string]map[string]struct{}{},
	}

	for _, word := range dictionary {
		abbr := getAbbr(word)
		m, ok := validator.abbrMap[abbr]
		if !ok {
			m = map[string]struct{}{}
		}

		m[word] = struct{}{}
		validator.abbrMap[abbr] = m
	}

	return validator
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	m, ok := this.abbrMap[getAbbr(word)]

	if !ok {
		return true
	}

	if len(m) > 1 {
		return false
	}

	_, inMap := m[word]

	return inMap
}
