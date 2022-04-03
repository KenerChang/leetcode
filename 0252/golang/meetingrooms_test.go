package meetingrooms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAttendMeetings(t *testing.T) {
	assert.False(t, canAttendMeetings([][]int{
		{0, 30}, {5, 10}, {15, 20},
	}))

	assert.False(t, canAttendMeetings([][]int{
		{0, 30}, {15, 20}, {5, 10},
	}))

	assert.True(t, canAttendMeetings([][]int{
		{7, 10}, {2, 4},
	}))

	assert.True(t, canAttendMeetings([][]int{
		{7, 10},
	}))

	assert.True(t, canAttendMeetings([][]int{}))
}
