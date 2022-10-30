package nestedlistweightsum

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

func depthSum(nestedList []*NestedInteger) int {
	var result int
	for _, nestedInt := range nestedList {
		result += dfs(nestedInt, 1)
	}

	return result
}

func dfs(nestedInt *NestedInteger, level int) int {
	// if current NestedInteger is a int
	// return nestedInt.GetInteger() * level
	if nestedInt.IsInteger() {
		return nestedInt.GetInteger() * level
	}

	// else it is a list
	// result = sum of its nested elements
	var result int
	for _, integer := range nestedInt.GetList() {
		result += dfs(integer, level+1)
	}

	return result
}
