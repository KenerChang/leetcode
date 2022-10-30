package nestedlistweightsumii

type NestedInteger struct{}

func (n NestedInteger) IsInteger() bool {
	return false
}

func (n NestedInteger) GetInteger() int {
	return 0
}

func (n *NestedInteger) SetInteger(value int) {}

func (n *NestedInteger) Add(elem NestedInteger) {}

func (n NestedInteger) GetList() []*NestedInteger {
	return []*NestedInteger{}
}

func count(integer *NestedInteger, maxDepth, level int) int {
	if integer.IsInteger() {
		weight := maxDepth - level + 1

		return integer.GetInteger() * weight
	}

	var result int
	for _, nestInteger := range integer.GetList() {
		result += count(nestInteger, maxDepth, level+1)
	}

	return result
}

func maxDepthOf(integer *NestedInteger, level int) int {
	if integer.IsInteger() {
		return level
	}

	var maxDepth int = level
	for _, nestedInteger := range integer.GetList() {
		maxDepth = max(maxDepth, maxDepthOf(nestedInteger, level+1))
	}

	return maxDepth
}

func depthSumInverse(nestedList []*NestedInteger) int {
	// first find maxDepth
	var maxDepth int
	for _, integer := range nestedList {
		maxDepth = max(maxDepthOf(integer, 1), maxDepth)
	}

	var result int
	for _, integer := range nestedList {
		result += count(integer, maxDepth, 1)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
