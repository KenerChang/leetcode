package insertdeletegetrandomoone

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	numsMap map[int]int
	nums    []int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	return RandomizedSet{
		numsMap: map[int]int{},
		nums:    []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.numsMap[val]

	if !ok {
		this.numsMap[val] = len(this.nums)
		this.nums = append(this.nums, val)
	}

	return !ok
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.numsMap[val]

	if ok {
		last := this.nums[len(this.nums)-1]

		this.numsMap[last] = idx
		this.nums[idx] = last

		this.nums = this.nums[0 : len(this.nums)-1]

		delete(this.numsMap, val)
	}

	return ok
}

func (this *RandomizedSet) GetRandom() int {

	idx := rand.Intn(len(this.nums))

	return this.nums[idx]
}
