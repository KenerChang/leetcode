package excelsheetcolumnnumber

// import (
// 	"fmt"
// )

func titleToNumber(s string) int {
	base := 1
	var result, digit int
	var char byte
	for i := len(s) - 1; i >= 0; i-- {
		char = s[i]
		digit = toNumber(char)
		result += digit * base
		base *= 26
	}
	return result
}

func toNumber(char byte) int {
	return int(char) - 64
}
