package maximumxoroftwonumbersinanarray

type bitwiseTrie struct {
	children []*bitwiseTrie
}

func (t *bitwiseTrie) Put(bits []int) {
	curr := t

	for _, bit := range bits {
		if curr.children[bit] == nil {
			curr.children[bit] = bitwiseTrieConstructor()
		}

		curr = curr.children[bit]
	}
}

func (t *bitwiseTrie) MaxXORMatch(bits []int) int {
	curr := t
	var maxXOR int
	for _, bit := range bits {
		if curr.children[1-bit] == nil && curr.children[bit] == nil {
			return 0
		}

		if curr.children[1-bit] != nil {
			// go opposite bit
			maxXOR = maxXOR | 1
			curr = curr.children[1-bit]
		} else {
			curr = curr.children[bit]
		}

		maxXOR = maxXOR << 1
	}

	maxXOR = maxXOR >> 1

	return maxXOR
}

func bitwiseTrieConstructor() *bitwiseTrie {
	return &bitwiseTrie{
		children: make([]*bitwiseTrie, 2),
	}
}

func bitsLength(num int) int {
	var count int
	for num > 0 {
		count++
		num = num >> 1
	}

	return count
}

func findMaximumXOR(nums []int) int {
	// // find max num in nums, and get how many bits used
	maxVal := nums[0]
	for i := 0; i < len(nums); i++ {
		maxVal = max(maxVal, nums[i])
	}
	maxBitLength := bitsLength(maxVal)

	// transform each num in nums to bits with padding
	numsBits := make([][]int, len(nums))
	for i, num := range nums {
		numsBits[i] = toBits(num, maxBitLength)
	}

	// for num in nums, put each num to the bitwise trie & count max xor value of this num (numMaxXor)
	// in the last of iteration set maxXor = max(maxXor, numMaxXor)
	var maxXOR int
	root := bitwiseTrieConstructor()
	for _, numBits := range numsBits {
		root.Put(numBits)

		numMaxXOR := root.MaxXORMatch(numBits)

		maxXOR = max(maxXOR, numMaxXOR)
	}

	return maxXOR
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func toBits(num int, L int) []int {
	// fmt.Printf("num: %d, L: %d\n", num, L)
	bits := make([]int, L)
	for i := L - 1; i >= 0; i-- {
		bits[L-i-1] = num >> i & 1
	}
	return bits
}
