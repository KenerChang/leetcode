package rangesumquerymutable

type NumArray struct {
	bit  []int
	nums []int
}

func Constructor(nums []int) NumArray {
	// construct a binary indexed tree (Fenwick tree)
	bit := make([]int, len(nums)+1)

	for i := 1; i < len(bit); i++ {
		bit[i] += nums[i-1]

		// propagate to next node
		parentIdx := i + lsb(i)
		if parentIdx < len(bit) {
			bit[parentIdx] += bit[i]
		}
	}

	// fmt.Printf("bit: %v\n", bit)

	return NumArray{
		bit:  bit,
		nums: nums,
	}
}

func (this *NumArray) Update(index int, val int) {
	bitIdx := index + 1

	diff := val - this.nums[index]
	this.nums[index] = val

	// propogate
	for bitIdx < len(this.bit) {
		this.bit[bitIdx] += diff

		bitIdx = bitIdx + lsb(bitIdx)
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.getSumOfIndex(right+1) - this.getSumOfIndex(left)
}

func (this *NumArray) getSumOfIndex(bitIdx int) int {
	var result int
	for bitIdx > 0 {
		result += this.bit[bitIdx]
		bitIdx -= lsb(bitIdx)
	}

	// fmt.Printf("bitIdx: %d, result: %d", bitIdx, result)

	return result
}

func lsb(i int) int {
	return i & -i
}
