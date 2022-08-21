package twosumiii

type TwoSum struct {
	nums  map[int]int
	cache map[int]bool
}

func Constructor() TwoSum {
	return TwoSum{
		nums:  make(map[int]int),
		cache: make(map[int]bool),
	}
}

func (ts *TwoSum) Add(number int) {
	ts.nums[number]++
}

func (ts *TwoSum) Find(value int) bool {
	// check if value is in cache
	if _, ok := ts.cache[value]; ok {
		return true
	}

	result := ts.twoSum(value)
	if result {
		ts.cache[value] = true
	}

	return result
}

func (ts *TwoSum) twoSum(value int) bool {
	for k, _ := range ts.nums {
		target := value - k
		if remain, ok := ts.nums[value-k]; ok {
			if target == k {
				if remain >= 2 {
					return true
				}
			} else {
				if remain >= 1 {
					return true
				}
			}
		}
	}
	return false
}
