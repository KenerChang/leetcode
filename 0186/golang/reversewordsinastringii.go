package reversewordsinastringii

func reverseWords(s []byte) {
	swap(s)

	head, ptr := 0, 0
	for ptr < len(s) {
		if s[ptr] != ' ' {
			ptr++
		} else {
			swap(s[head:ptr])
			ptr++
			head = ptr
		}
	}

	swap(s[head:ptr])
}

func swap(s []byte) {
	head, last := 0, len(s)-1

	for head < last {
		s[head], s[last] = s[last], s[head]
		head++
		last--
	}
}
