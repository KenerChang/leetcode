package golang

import (
	"math/rand"
	"time"
)

var (
	r *rand.Rand
)

func rand7() int {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	return r.Intn(7) + 1
}

func rand10() int {

	row := rand7()
	col := rand7()

	for (row-1)*7+col > 40 {
		row = rand7()
		col = rand7()
	}

	return ((row-1)*7+col)%10 + 1
}
