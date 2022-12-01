package decodestring

import (
	"strconv"
)

func decodeString(s string) string {
	// We can use stack to keep encoded patterns
	stack := [][]byte{}
	freqs := []int{}
	freqBytes := []byte{}

	result := []byte{}
	for _, c := range s {
		if c >= '0' && c <= '9' {
			freqBytes = append(freqBytes, byte(c))
		} else {
			if c == ']' {
				last := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				freq := freqs[len(freqs)-1]
				freqs = freqs[:len(freqs)-1]

				decoded := []byte{}
				for i := 0; i < freq; i++ {
					decoded = append(decoded, last...)
				}

				if len(stack) == 0 {
					// output to result
					result = append(result, decoded...)
				} else {
					// put it to the last group in the stack
					stack[len(stack)-1] = append(stack[len(stack)-1], decoded...)
				}

			} else if c == '[' {
				stack = append(stack, []byte{})

				freq, _ := strconv.Atoi(string(freqBytes))
				freqBytes = []byte{}

				freqs = append(freqs, freq)
			} else {
				if len(stack) == 0 {
					// pattern like a, output to result
					result = append(result, byte(c))
				} else {
					stack[len(stack)-1] = append(stack[len(stack)-1], byte(c))
				}
			}
		}

	}
	return string(result)
}
