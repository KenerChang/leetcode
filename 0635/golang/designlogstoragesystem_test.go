package designlogstoragesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogSystemI(t *testing.T) {
	ls := Constructor()

	ls.Put(1, "2017:01:01:23:59:59")
	ls.Put(2, "2017:01:01:22:59:59")
	ls.Put(3, "2016:01:01:00:00:00")

	assert.Equal(t, []int{3, 2, 1}, ls.Retrieve("2016:01:01:01:01:01", "2017:01:01:23:00:00", "Year"))
	assert.Equal(t, []int{2, 1}, ls.Retrieve("2016:01:01:01:01:01", "2017:01:01:23:00:00", "Hour"))
}

func TestLogSystemII(t *testing.T) {
	ls := Constructor()

	ls.Put(1, "2017:01:01:23:59:59")
	ls.Put(2, "2017:01:02:23:59:59")

	assert.Equal(t, []int{1}, ls.Retrieve("2017:01:01:23:59:58", "2017:01:02:23:59:58", "Second"))
}

func TestLogSystemIII(t *testing.T) {
	ls := Constructor()

	ls.Put(1, "2017:01:01:23:59:59")

	assert.Equal(t, []int{1}, ls.Retrieve("2017:01:01:23:59:58", "2017:01:02:23:59:58", "Second"))
}

func TestLogSystemIV(t *testing.T) {
	ls := Constructor()

	ls.Put(1, "2005:01:05:22:16:15")
	ls.Put(2, "2003:12:12:20:30:51")

	assert.Equal(t, []int{}, ls.Retrieve("2005:07:10:17:43:43", "2007:02:18:10:22:52", "Month"))
}
