package nextclosesttime

import "time"

func isValid(t time.Time, digits []bool) bool {
	timeString := t.Format("15:04")

	for i := range timeString {
		if timeString[i] == ':' {
			continue
		}

		if !digits[timeString[i]-'0'] {
			return false
		}
	}

	return true
}

func nextClosestTime(timeString string) string {
	digits := make([]bool, 10)
	count := 0
	for i := range timeString {
		if timeString[i] == ':' {
			continue
		}

		if !digits[timeString[i]-'0'] {
			count++
			digits[timeString[i]-'0'] = true
		}
	}

	// pattern like "00:00"
	if count == 1 {
		return timeString
	}

	t, _ := time.Parse("15:04", timeString)
	t = t.Add(1 * time.Minute)

	for !isValid(t, digits) {
		t = t.Add(1 * time.Minute)
	}

	return t.Format("15:04")
}
