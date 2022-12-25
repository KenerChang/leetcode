package designphonedirectory

import "container/list"

type PhoneDirectory struct {
	AvailableNums *list.List
	Used          map[int]*list.Element
}

func Constructor(maxNumbers int) PhoneDirectory {
	d := PhoneDirectory{
		Used: map[int]*list.Element{},
	}

	l := list.New()
	for i := 0; i < maxNumbers; i++ {
		l.PushBack(i)
	}
	d.AvailableNums = l

	return d
}

func (d *PhoneDirectory) Get() int {
	top := d.AvailableNums.Front()
	if top == nil {
		return -1
	}
	d.AvailableNums.Remove(top)

	d.Used[top.Value.(int)] = top

	return top.Value.(int)
}

func (d *PhoneDirectory) Check(number int) bool {
	_, used := d.Used[number]
	return !used
}

func (d *PhoneDirectory) Release(number int) {
	e, used := d.Used[number]
	if used {
		delete(d.Used, number)
		d.AvailableNums.PushBack(e.Value.(int))
	}
}
