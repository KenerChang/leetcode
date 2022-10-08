package numberOfCombinations

const (
	m int = 1e9 + 7
)

func buildCommonPrefix(num string) [][]int {
	commons := make([][]int, len(num))
	for i := len(num) - 1; i >= 0; i-- {
		commons[i] = make([]int, len(num))
		for j := len(num) - 1; j >= 0; j-- {
			if num[i] == num[j] {
				if i == len(num)-1 || j == len(num)-1 {
					commons[i][j] = 1
				} else {
					commons[i][j] = commons[i+1][j+1] + 1
				}
			}
		}
	}

	return commons
}

func compare(i, j, len int, commons [][]int, num string) bool {
	common := commons[i][j]

	if common >= len {
		return true
	}

	if num[i+common] < num[j+common] {
		return true
	}

	return false
}

func mod(a, b int) int {
	return (a + b) % m
}

func numberOfCombinations(num string) int {
	// dynmaic programming
	// prerequisites:
	// 1. how to compare two string in O(1)
	// 2. how to fetch prefix sum quickly

	if num[0] == '0' {
		return 0
	}

	// dp[i][j] means last number is from i to j
	// transition function: dp[i][j] = dp[i-length+1][i-1] + dp[i-length+2][i-1] + ...+  dp[i-1][i-1]

	// To solve prerequisties 1, we can use a 2D array to store most common prefix
	// of all substrings starting from i and j
	commons := buildCommonPrefix(num)

	// use an array to store prefix sum of dp[0][j] + dp[1][j] ... + dp[i][j]
	prefixSums := make([][]int, len(num))
	for i := 0; i < len(num); i++ {
		prefixSums[i] = make([]int, len(num))
	}

	for j := 0; j < len(num); j++ {
		prefixSums[0][j] = 1
	}

	for i := 1; i < len(num); i++ {
		if num[i] == '0' {
			prefixSums[i] = prefixSums[i-1]
			continue
		}

		for j := i; j < len(num); j++ {
			len := j - i + 1
			prevStart := i - 1 - (len - 1)
			count := 0

			if prevStart < 0 {
				count = prefixSums[i-1][i-1]
			} else {
				count = mod(prefixSums[i-1][i-1], -prefixSums[prevStart][i-1]) //= dp[prevStart+1][i-1] ... dp[i-1][i-1]

				if compare(prevStart, i, len, commons, num) {
					var cnt int
					if prevStart == 0 {
						cnt = prefixSums[prevStart][i-1]
					} else {
						cnt = mod(prefixSums[prevStart][i-1], -prefixSums[prevStart-1][i-1])
					}
					count = mod(count, cnt)
				}
			}

			prefixSums[i][j] = mod(prefixSums[i-1][j], count) //updating prefix sum
		}
	}
	return prefixSums[len(num)-1][len(num)-1]
}
