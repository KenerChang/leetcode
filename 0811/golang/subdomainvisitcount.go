package subdomainvisitcount

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	// input format "<num> d1.d2.d3" or "<num> d1.d2"

	// use a map to keep domain visited counts
	domainCounts := map[string]int{}

	for _, cpdomain := range cpdomains {
		// parse visited count
		tokens := strings.Split(cpdomain, " ")
		count, _ := strconv.Atoi(tokens[0])

		// parse domains
		domains := parseDomains(tokens[1])

		for _, domain := range domains {
			domainCounts[domain] += count
		}
	}

	// generate output
	results := []string{}
	for domain, count := range domainCounts {
		results = append(results, fmt.Sprintf("%d %s", count, domain))
	}
	return results
}

func parseDomains(cpdomain string) []string {
	ds := strings.Split(cpdomain, ".")

	// iterate domains in reverse order
	var domain string = ds[len(ds)-1]
	results := []string{ds[len(ds)-1]}
	for i := len(ds) - 2; i >= 0; i-- {
		domain = fmt.Sprintf("%s.%s", ds[i], domain)

		results = append(results, domain)
	}
	return results
}
