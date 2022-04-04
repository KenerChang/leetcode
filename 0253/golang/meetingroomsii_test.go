package meetingroomsii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAttendMeetings(t *testing.T) {
	assert.Equal(t, 2, minMeetingRooms([][]int{
		{0, 30}, {5, 10}, {15, 20},
	}))

	assert.Equal(t, 1, minMeetingRooms([][]int{
		{7, 10}, {2, 4},
	}))

}
