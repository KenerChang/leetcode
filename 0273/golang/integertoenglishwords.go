package integertoenglishwords

import (
	"fmt"
	"math"
	"strings"
)

var (
	BILLION                     = int(math.Pow10(9))
	MILLION                     = int(math.Pow10(6))
	THOUSAND                    = int(math.Pow10(3))
	ZERO                        = 0
	HUNDRED                     = int(math.Pow10(2))
	orders       []int          = []int{BILLION, MILLION, THOUSAND}
	numStringMap map[int]string = map[int]string{
		1: "One",
		2: "Two",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
		3: "Three",
	}
	tensNumStringMap map[int]string = map[int]string{
		2: "Twenty",
		4: "Forty",
		5: "Fifty",
		6: "Sixty",
		7: "Seventy",
		8: "Eighty",
		9: "Ninety",
		3: "Thirty",
	}
	eleToNineteenMap map[int]string = map[int]string{
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
	}
	orderStringMap map[int]string = map[int]string{
		BILLION:  "Billion",
		MILLION:  "Million",
		THOUSAND: "Thousand",
	}
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	// for order in orders
	remain := num
	result := ""
	for _, order := range orders {
		// if num / order > 0, tranform to string
		quotient := remain / order
		if quotient > 0 {
			result = fmt.Sprintf("%s %s", result, transformIntToString(quotient, order))
		}

		// remain = remain % order
		remain = remain % order
	}

	if remain > 0 {
		result = fmt.Sprintf("%s %s", result, transformIntToString(remain, ZERO))
	}

	return strings.TrimSpace(result)
}

func transformIntToString(num int, base int) string {
	result := ""

	quotient := num / HUNDRED
	// handle hundreds
	if quotient > 0 {
		result = fmt.Sprintf("%s Hundred", numStringMap[quotient])
	}

	// handle 2 digits and one digit
	num = num % HUNDRED

	if num > 0 {
		// if num < 10, suffix string to result
		if num < 10 {
			result = fmt.Sprintf("%s %s", result, numStringMap[num])
		} else if num < 20 && num >= 10 {
			// else if num<20 and num >= 10, and ten, eleven, ..., nineteen to result
			result = fmt.Sprintf("%s %s", result, eleToNineteenMap[num])
		} else {
			// else get tens then get one digit
			result = fmt.Sprintf("%s %s", result, tensNumStringMap[num/10])
			if num%10 > 0 {
				result = fmt.Sprintf("%s %s", result, numStringMap[num%10])
			}
		}
	}

	if base != ZERO {
		result = fmt.Sprintf("%s %s", result, orderStringMap[base])
	}

	return strings.TrimSpace(result)
}
