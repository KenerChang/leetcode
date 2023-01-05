package largesttimeforgivendigits

import "fmt"

func largestTimeFromDigits(arr []int) string {
	// collect digits
	digitFrequencies := make([]int, 10)
	for _, digit := range arr {
		digitFrequencies[digit]++
	}

	var tdHour, udHour, tdMin, udMin int
	for i := 2; i >= 0; i-- {
		if digitFrequencies[i] <= 0 {
			continue
		}

		tdHour = i
		digitFrequencies[i]--
		var j int = 9
		if i == 2 {
			j = 3
		}

		for j >= 0 {
			if digitFrequencies[j] <= 0 {
				j--
				continue
			}

			udHour = j
			digitFrequencies[j]--

			for k := 5; k >= 0; k-- {
				if digitFrequencies[k] <= 0 {
					continue
				}

				tdMin = k
				digitFrequencies[k]--

				for m := 9; m >= 0; m-- {
					if digitFrequencies[m] <= 0 {
						continue
					}

					udMin = m
					return fmt.Sprintf("%d%d:%d%d", tdHour, udHour, tdMin, udMin)
				}

				digitFrequencies[k]++
			}

			digitFrequencies[j]++
			j--
		}

		digitFrequencies[i]++
	}

	return ""
}
