package flattennestedlistiterator

type NestedInteger struct{}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return []*NestedInteger{}
}

func flatten(nested *NestedInteger) []int {
	if nested.IsInteger() {
		return []int{nested.GetInteger()}
	}

	var res []int
	for _, v := range nested.GetList() {
		res = append(res, flatten(v)...)
	}
	return res
}

type NestedIterator struct {
	// nesteds []NestedInteger
	Nums []int
}

func Constructor(nestedList []*NestedInteger) NestedIterator {
	// flatten nested list
	nums := []int{}
	for _, v := range nestedList {
		nums = append(nums, flatten(v)...)
	}
	return NestedIterator{
		Nums: nums,
	}
}

func (i *NestedIterator) Next() int {
	result := i.Nums[0]

	i.Nums = i.Nums[1:]
	return result
}

func (i *NestedIterator) HasNext() bool {
	return len(i.Nums) > 0
}
