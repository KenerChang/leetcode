package rangesumquery2d

type NumMatrix struct {
	bits   [][]int
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	// build bit for each row
	bits := make([][]int, len(matrix))

	for i, row := range matrix {
		bit := make([]int, len(row)+1)

		for j, val := range row {
			bitIdx := j + 1
			bit[bitIdx] += val

			nextNodeIex := bitIdx + lsb(bitIdx)
			if nextNodeIex < len(bit) {
				bit[nextNodeIex] += bit[bitIdx]
			}
		}

		bits[i] = bit
	}

	// fmt.Printf("bits: %+v\n", bits)

	return NumMatrix{
		matrix: matrix,
		bits:   bits,
	}
}

func (this *NumMatrix) Update(row int, col int, val int) {
	diff := val - this.matrix[row][col]
	this.matrix[row][col] = val

	bit := this.bits[row]
	bitIdx := col + 1

	for bitIdx < len(bit) {
		bit[bitIdx] += diff
		bitIdx += lsb(bitIdx)
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var result int

	for i := row1; i <= row2; i++ {
		bit := this.bits[i]
		// fmt.Printf("row: %d, col2: %d, col1: %d\n", i, this.doQuery(bit, col2+1), this.doQuery(bit, col1))
		result += (this.doQuery(bit, col2+1) - this.doQuery(bit, col1))
	}

	return result
}

func (this *NumMatrix) doQuery(bit []int, bitIdx int) int {
	var result int
	for bitIdx > 0 {
		result += bit[bitIdx]
		bitIdx -= lsb(bitIdx)
	}

	return result
}

func lsb(i int) int {
	return i & -i
}
