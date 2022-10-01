package peekingiterator

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type Iterator struct {
	nums []int
	pos  int
}

func (i *Iterator) hasNext() bool {
	if len(i.nums) > 0 && i.pos < len(i.nums)-1 {
		return true
	}
	return false
}

func (i *Iterator) next() int {
	if i.hasNext() {
		i.pos++
		return i.nums[i.pos]
	}

	return -1
}

type PeekingIterator struct {
	iter    *Iterator
	peekNum *int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter}
}

func (pi *PeekingIterator) hasNext() bool {
	if pi.peekNum != nil || pi.iter.hasNext() {
		return true
	}

	return false
}

func (pi *PeekingIterator) next() int {
	if pi.peekNum != nil {
		num := *pi.peekNum
		pi.peekNum = nil
		return num
	}
	return pi.iter.next()
}

func (pi *PeekingIterator) peek() int {
	if pi.peekNum != nil {
		return *pi.peekNum
	}

	num := pi.iter.next()
	pi.peekNum = &num
	return num
}
