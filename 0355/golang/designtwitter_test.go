package designtwitter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwitterI(t *testing.T) {
	twitter := Constructor()

	twitter.PostTweet(1, 5)

	assert.Equal(t, []int{5}, twitter.GetNewsFeed(1))

	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)

	assert.Equal(t, []int{6, 5}, twitter.GetNewsFeed(1))

	twitter.Unfollow(1, 2)

	assert.Equal(t, []int{5}, twitter.GetNewsFeed(1))
}

func TestTwittedII(t *testing.T) {
	twitter := Constructor()

	twitter.PostTweet(1, 1)

	assert.Equal(t, []int{1}, twitter.GetNewsFeed(1))

	twitter.Follow(2, 1)

	assert.Equal(t, []int{1}, twitter.GetNewsFeed(2))

	twitter.Unfollow(2, 1)

	assert.Equal(t, []int{}, twitter.GetNewsFeed(2))
}
