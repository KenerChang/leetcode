package encodeanddecodestrings

import (
	"fmt"
	"strconv"
)

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	result := ""

	// iterate strs, add length as prefix of the str
	for _, str := range strs {
		result = fmt.Sprintf("%s%d$%s", result, len(str), str)
	}

	return result
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	results := []string{}

	ptr := 0
	var str string
	// while not reach end of the string
	for ptr < len(strs) {
		// get lentgh of next string
		// move ptr

		str, ptr = codec.parseNextStr(strs, ptr)
		results = append(results, str)
	}
	return results
}

func (codec *Codec) parseNextStr(strs string, ptr int) (string, int) {
	bytes := []byte{}
	for strs[ptr] != '$' {
		bytes = append(bytes, strs[ptr])
		ptr += 1
	}

	ptr++
	strLen, _ := strconv.Atoi(string(bytes))
	if strLen == 0 {
		return "", ptr
	}

	nextHead := ptr + strLen
	return strs[ptr:nextHead], nextHead
}
