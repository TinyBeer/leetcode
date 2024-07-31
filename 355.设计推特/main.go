package main

import "sort"

/*
设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近 10 条推文。

实现 Twitter 类：

Twitter() 初始化简易版推特对象
void postTweet(int userId, int tweetId) 根据给定的 tweetId 和 userId 创建一条新推文。每次调用此函数都会使用一个不同的 tweetId 。
List<Integer> getNewsFeed(int userId) 检索当前用户新闻推送中最近  10 条推文的 ID 。新闻推送中的每一项都必须是由用户关注的人或者是用户自己发布的推文。推文必须 按照时间顺序由最近到最远排序 。
void follow(int followerId, int followeeId) ID 为 followerId 的用户开始关注 ID 为 followeeId 的用户。
void unfollow(int followerId, int followeeId) ID 为 followerId 的用户不再关注 ID 为 followeeId 的用户。
*/

type Twitter struct {
	userMap map[int]*User
}

type User struct {
	userId    int
	followees map[int]bool
	tweets    []*Tweet
}

type Tweet struct {
	tweetId int
	time    int
	userId  int
}

// 推特数目，用于时间排序
var tweetCount int

func Constructor() Twitter {
	return Twitter{userMap: make(map[int]*User)}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	// 新建tweet  将自己设置为关注

	// 如果map 中不存在需要新建，因为User 类中存在map 和 slice
	if _, ok := t.userMap[userId]; !ok {
		t.userMap[userId] = &User{userId: userId, tweets: make([]*Tweet, 0), followees: make(map[int]bool)}
		tweet := &Tweet{tweetId, tweetCount, userId}
		t.userMap[userId].tweets = append(t.userMap[userId].tweets, tweet)
		t.userMap[userId].followees[userId] = true
	} else {
		tweet := &Tweet{tweetId, tweetCount, userId}
		t.userMap[userId].tweets = append(t.userMap[userId].tweets, tweet)

	}

	// 将tweetId 和时间做一个新增
	tweetCount++
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	// 如果关注人不存在则新建
	if _, ok := t.userMap[followerId]; !ok {
		t.userMap[followerId] = &User{
			userId:    followerId,
			followees: make(map[int]bool),
		}
		// 每次新建user的时候 将自己加入自己关注
		t.userMap[followerId].followees[followerId] = true
	}

	// 如果被关注人不存在则新建
	if _, ok := t.userMap[followeeId]; !ok {
		t.userMap[followeeId] = &User{
			userId:    followeeId,
			followees: make(map[int]bool),
		}
		// 每次新建user的时候 将自己加入自己关注
		t.userMap[followeeId].followees[followeeId] = true
	}
	t.userMap[followerId].followees[followeeId] = true
}

// 形参上的Id 在使用数据结构的时候一般使用map查找
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := t.userMap[followerId]; ok {
		delete(t.userMap[followerId].followees, followeeId)
	}
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	resTop10 := []int{}
	tweeters := []*Tweet{}
	if _, ok := t.userMap[userId]; ok {
		for followeeId := range t.userMap[userId].followees {
			tweeters = append(tweeters, t.userMap[followeeId].tweets...)
		}
	}
	sort.Slice(tweeters, func(i, j int) bool {
		if tweeters[i].time > tweeters[j].time {
			return true
		}
		return false
	})

	for i := 0; i < len(tweeters) && i < 10; i++ {
		resTop10 = append(resTop10, tweeters[i].tweetId)
	}
	// for idx, tweet := range tweeters {
	//     if idx > 9 {
	// 		break
	// 	}
	// 	resTop10 = append(resTop10, tweet.tweetId)

	// }
	return resTop10
}

// type User struct {
// 	ID        int
// 	Tweets    []*Tweet
// 	Followees map[int]struct{}
// }

// type Tweet struct {
// 	ID       int
// 	Sequence int
// }

// type Twitter struct {
// 	Sequence int
// 	Users    map[int]*User
// }

// func Constructor() Twitter {
// 	return Twitter{
// 		Users: map[int]*User{},
// 	}
// }

// func (t *Twitter) getUser(id int) *User {
// 	if u, ok := t.Users[id]; ok {
// 		return u
// 	}
// 	u := &User{
// 		ID:        id,
// 		Followees: map[int]struct{}{},
// 	}
// 	t.Users[id] = u
// 	return u
// }

// func (this *Twitter) PostTweet(userId int, tweetId int) {
// 	u := this.getUser(userId)
// 	u.Tweets = append(u.Tweets, &Tweet{
// 		ID:       tweetId,
// 		Sequence: this.Sequence,
// 	})
// 	this.Sequence++

// 	if len(u.Tweets) > 10 {
// 		u.Tweets = u.Tweets[len(u.Tweets)-10:]
// 	}
// }

// func (this *Twitter) GetNewsFeed(userId int) []int {
// 	u := this.getUser(userId)
// 	var tweetList []*Tweet
// 	tweetList = append(tweetList, u.Tweets...)
// 	for followeeID := range u.Followees {
// 		followee := this.getUser(followeeID)
// 		tweetList = append(tweetList, followee.Tweets...)
// 	}
// 	sort.Slice(tweetList, func(i, j int) bool {
// 		return tweetList[i].Sequence > tweetList[j].Sequence
// 	})
// 	list := make([]int, 0, 10)
// 	for i := 0; i < 10 && i < len(tweetList); i++ {
// 		list = append(list, tweetList[i].ID)
// 	}
// 	return list
// }

// func (this *Twitter) Follow(followerId int, followeeId int) {
// 	follower := this.getUser(followerId)
// 	follower.Followees[followeeId] = struct{}{}
// }

// func (this *Twitter) Unfollow(followerId int, followeeId int) {
// 	follower := this.getUser(followerId)
// 	delete(follower.Followees, followeeId)
// }

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
