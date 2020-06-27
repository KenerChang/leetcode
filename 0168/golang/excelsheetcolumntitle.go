package excelsheetcolumntitle

// import (
// 	"fmt"
// )

func convertToTitle(n int) string {
	if n <= 0 {
		return ""
	}

	base := 26

	mappingTable := []string{
		"Z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y",
	}

	var resultStr string
	var quotient, remainder int = n, 0

	for quotient > 0 {
		remainder = quotient % base
		quotient = quotient / base
		if remainder == 0 {
			quotient--
		}

		resultStr = mappingTable[remainder] + resultStr
	}

	return resultStr
}
