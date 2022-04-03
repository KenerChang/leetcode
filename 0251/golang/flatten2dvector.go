package flatten2dvector

type Vector2D struct {
	vector   [][]int
	rowPtr   int
	colPtr   int
	iterated bool
}

func Constructor(vec [][]int) Vector2D {
	v := Vector2D{
		vector: vec,
	}

	if len(vec) == 0 {
		v.iterated = true
		return v
	}

	if len(vec[0]) == 0 {
		v.movePtrs()
	}

	return v
}

func (this *Vector2D) Next() int {
	result := this.vector[this.rowPtr][this.colPtr]

	// move ptr
	this.movePtrs()

	return result
}

func (this *Vector2D) movePtrs() {
	// if colPtr does not reach end of this row, colPtr++
	this.colPtr++
	if this.colPtr < len(this.vector[this.rowPtr]) {
		return
	}

	// while colPtr reaches end of this row, move to next row
	for this.rowPtr < len(this.vector) && this.colPtr >= len(this.vector[this.rowPtr]) {
		this.colPtr = 0
		this.rowPtr++
	}

	// if rows are iterated, set iterated to true
	if this.rowPtr >= len(this.vector) {
		this.iterated = true
		return
	}
}

func (this *Vector2D) HasNext() bool {
	return !this.iterated
}
