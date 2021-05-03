package oneeditdistance

func isOneEditDistance(s string, t string) bool {
	return levEqual1(s, t, 0)
}

func levEqual1(s string, t string, accu int) bool {
	// fmt.Printf("s: %s, t: %s, accu: %d\n", s, t, accu)
	if accu > 1 {
		return false
	}

	if len(s) == 0 {
		return len(t)+accu == 1
	}

	if len(t) == 0 {
		return len(s)+accu == 1
	}

	if s[0] == t[0] {
		return levEqual1(s[1:], t[1:], accu)
	}

	return levEqual1(s[1:], t, accu+1) || levEqual1(s, t[1:], accu+1) || levEqual1(s[1:], t[1:], accu+1)
}
