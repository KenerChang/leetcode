package rangesumqueryimmutable

type NumArray struct {
	nums  []int
	cache []int
}

func Constructor(nums []int) NumArray {
	a := NumArray{
		nums:  nums,
		cache: make([]int, len(nums)+1),
	}
	a.cache[0] = 0

	for i := 1; i <= len(nums); i++ {
		a.cache[i] += a.cache[i-1] + nums[i-1]
	}

	return a
}

func (a *NumArray) SumRange(left int, right int) int {
	return a.cache[right+1] - a.cache[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
