package readncharactersgivenread4

func read4Impl(input string) func([]byte) int {
	var read int
	return func(buf []byte) int {
		remain := len(input) - read
		end := min(4, remain)
		copy(buf, input[read:read+end])
		read += end
		return end
	}
}

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		// each time we read from read4
		// handle 2 situations:
		// 1. read4 reaches EOF
		// 2. read function reaches n

		// read from read4
		buf4 := make([]byte, 4)
		remain := n
		read := 0
		for remain > 0 {
			need := min(remain, read4(buf4))

			if need == 0 {
				// file reaches EOF
				break
			}

			// copy to buf
			copy(buf[read:], buf4[:need])

			read += need
			remain -= need
		}

		return read
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
