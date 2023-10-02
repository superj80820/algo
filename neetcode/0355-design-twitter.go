import "container/heap"

var FeedMaxCount int = 10

type TweetInfoPQ []*PostInfo

func (pq TweetInfoPQ) Len() int           { return len(pq) }
func (pq TweetInfoPQ) Less(i, j int) bool { return pq[i].CreateTime > pq[j].CreateTime }
func (pq TweetInfoPQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *TweetInfoPQ) Push(x any) {
	*pq = append(*pq, x.(*PostInfo))
}
func (pq *TweetInfoPQ) Pop() any {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return x
}

type Twitter struct {
	FollowMap  map[int]map[int]bool
	TweetMap   map[int][]*PostInfo
	TweetCount int
}

type PostInfo struct {
	TweetId    int
	CreateTime int
}

func Constructor() Twitter {
	return Twitter{
		FollowMap: make(map[int]map[int]bool),
		TweetMap:  make(map[int][]*PostInfo),
	}
}

// time complexity: O(1)
// space complexity: O(1)
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.TweetCount++
	this.TweetMap[userId] = append(this.TweetMap[userId], &PostInfo{
		TweetId:    tweetId,
		CreateTime: this.TweetCount,
	})
}

// time complexity: O(n)
// space complexity: O(1)
// `n` is follow count
func (this *Twitter) GetNewsFeed(userId int) []int {
	postUsers := []int{userId}
	if follwers, ok := this.FollowMap[userId]; ok {
		for follwer, _ := range follwers {
			postUsers = append(postUsers, follwer)
		}
	}

	pq := new(TweetInfoPQ)
	for i := 0; i < FeedMaxCount; i++ {
		for _, postUser := range postUsers {
			if posts, ok := this.TweetMap[postUser]; ok {
				idx := len(posts) - 1 - i
				if idx >= 0 {
					*pq = append(*pq, posts[idx])
				}
			}
		}
	}

	heap.Init(pq)

	var res []int
	for i := 0; i < FeedMaxCount; i++ {
		if pq.Len() == 0 {
			break
		}
		res = append(res, heap.Pop(pq).(*PostInfo).TweetId)
	}

	return res
}

// time complexity: O(1)
// space complexity: O(1)
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.FollowMap[followerId]; !ok {
		this.FollowMap[followerId] = make(map[int]bool)
	}
	this.FollowMap[followerId][followeeId] = true
}

// time complexity: O(1)
// space complexity: O(1)
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.FollowMap[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
