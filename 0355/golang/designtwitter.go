package designtwitter

import "container/list"

type tweet struct {
	ID     int
	UserID int
}

type twitterUser struct {
	ID        int
	Followees map[int]bool
}

func (u *twitterUser) Follow(followee *twitterUser) {
	u.Followees[followee.ID] = true
}

func (u *twitterUser) GetNewsFeed(tweets *list.List, n int) []int {
	var count int
	ele := tweets.Front()

	results := []int{}
	for ele != nil && count < n {
		t := ele.Value.(*tweet)

		if _, ok := u.Followees[t.UserID]; ok {
			results = append(results, t.ID)
			count++
		}

		ele = ele.Next()
	}
	return results
}

func (u *twitterUser) Unfollow(followee *twitterUser) {
	if followee.ID == u.ID {
		return
	}

	delete(u.Followees, followee.ID)
}

type Twitter struct {
	Tweets *list.List
	Users  map[int]*twitterUser
}

func Constructor() Twitter {
	return Twitter{
		Tweets: list.New(),
		Users:  map[int]*twitterUser{},
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := t.Users[userId]; !ok {
		user := newUser(userId)

		t.Users[userId] = user
	}

	t.Tweets.PushFront(&tweet{
		ID:     tweetId,
		UserID: userId,
	})
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	user, ok := t.Users[userId]
	if !ok {
		return []int{}
	}

	return user.GetNewsFeed(t.Tweets, 10)
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	follower, ok := t.Users[followerId]
	if !ok {
		follower = newUser(followerId)
		t.Users[followerId] = follower
	}

	followee, ok := t.Users[followeeId]
	if !ok {
		followee = newUser(followeeId)
		t.Users[followeeId] = followee
	}

	follower.Follow(followee)
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	follower, ok := t.Users[followerId]
	if !ok {
		return
	}

	followee, ok := t.Users[followeeId]
	if !ok {
		return
	}

	follower.Unfollow(followee)
}

func newUser(userId int) *twitterUser {
	return &twitterUser{
		ID: userId,
		Followees: map[int]bool{
			userId: true,
		},
	}
}
