package singlenumberii

func singleNumber(nums []int) int {
	var flag1, flag2 int
	for _, num := range nums {
		flag1 = flag1 ^ num&^flag2
		flag2 = flag2 ^ num&^flag1
	}
	return flag1
}
