package lexicographicalnumbers

type intTrie struct {
	children []*intTrie
}

func (t *intTrie) Put(n int) {
	digits := toDecimalDigits(n)

	// fmt.Printf("n: %d, digits: %v\n", n, digits)

	current := t
	for _, digit := range digits {
		if current.children[digit] == nil {
			current.children[digit] = trieConstructor()
		}

		current = current.children[digit]
	}
}

func (t *intTrie) Output(base int, nums []int) []int {
	for i := range t.children {
		if t.children[i] != nil {
			nums = append(nums, base+i)

			nums = t.children[i].Output((base+i)*10, nums)
		}
	}

	return nums
}

func lexicalOrder(n int) []int {
	root := trieConstructor()
	for i := 1; i <= n; i++ {
		root.Put(i)
	}

	return root.Output(0, []int{})
}

func toDecimalDigits(n int) []int {
	results := []int{}

	for n > 0 {
		results = append(results, n%10)

		n = n / 10
	}

	// swap
	var i, j int = 0, len(results) - 1
	for i < j {
		results[i], results[j] = results[j], results[i]
		i++
		j--
	}

	return results
}

func trieConstructor() *intTrie {
	return &intTrie{
		children: make([]*intTrie, 10),
	}
}
