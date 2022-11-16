package mycalendari

type MyCalendar struct {
	Books [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		Books: [][]int{},
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	// There 3 situations which will cause conflict
	// 1. exist a book b that b.end > start >= b.start
	// 2. exist a book b that start < b.end <= end
	if len(this.Books) == 0 {
		this.putBook(start, end, 0)
		return true
	}

	// for case 1, find the book which has the latest start time and is eariler than start
	lBook, idx := this.findLeft(start)
	if idx != -1 && lBook[1] > start {
		return false
	}

	// for case 2, find the book which has the eariles start time and is later than start
	if idx == len(this.Books)-1 {
		// no right book
		this.putBook(start, end, len(this.Books))
		return true
	}

	// check if these current book conflicts with these 2 books
	rBook := this.Books[idx+1]
	// fmt.Printf("start:%d, end: %d, rbook: %v\n", start, end, rBook)
	if end > rBook[0] {
		// fmt.Printf("end: %d, rBook[0]: %d\n", end, rBook[0])
		return false
	}

	// if not add it to the books
	this.putBook(start, end, idx+1)
	return true
}

func (this *MyCalendar) findLeft(start int) ([]int, int) {
	l, r := 0, len(this.Books)-1

	var resultIdx int = -1
	for l <= r {
		mid := (l + r) / 2

		if this.Books[mid][0] <= start {
			// search right

			resultIdx = mid
			l = mid + 1
		} else {
			// search left
			r = mid - 1
		}
	}

	if resultIdx == -1 {
		return []int{}, -1
	}

	return this.Books[resultIdx], resultIdx
}

func (this *MyCalendar) putBook(start, end, idx int) {
	newBooks := make([][]int, len(this.Books)+1)

	for i := 0; i < len(newBooks); i++ {
		if i == idx {
			newBooks[i] = []int{start, end}
		} else if i < idx {
			newBooks[i] = this.Books[i]
		} else {
			newBooks[i] = this.Books[i-1]
		}
	}

	this.Books = newBooks
}
