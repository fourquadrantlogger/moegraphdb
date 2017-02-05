package graphdb

// 关注他
func (UserArray RelateGraph) Like(vid, beliked uint) {
	UserArray.GetOrCreateUser(vid).Lock_likes.Lock()
	UserArray.GetOrCreateUser(vid).Likes[beliked] = true
	UserArray.GetOrCreateUser(vid).Lock_likes.Unlock()

	UserArray.GetOrCreateUser(beliked).Lock_fans.Lock()
	UserArray.GetOrCreateUser(beliked).Fans[vid] = true
	UserArray.GetOrCreateUser(beliked).Lock_fans.Unlock()
}

// 取消关注他
func (UserArray RelateGraph) DisLike(vid, beliked uint) {
	UserArray.GetOrCreateUser(vid).Lock_likes.Lock()
	delete(UserArray.GetOrCreateUser(vid).Likes, beliked)
	UserArray.GetOrCreateUser(vid).Lock_likes.Unlock()

	UserArray.GetOrCreateUser(beliked).Lock_fans.Lock()
	delete(UserArray.GetOrCreateUser(beliked).Fans, vid)
	UserArray.GetOrCreateUser(beliked).Lock_fans.Unlock()
}

// 互粉
func (UserArray RelateGraph) Makefriend(vid, beliked uint) {
	UserArray.GetOrCreateUser(vid).Lock_likes.Lock()
	UserArray.GetOrCreateUser(beliked).Lock_fans.Lock()
	UserArray.GetOrCreateUser(vid).Likes[beliked] = true
	UserArray.GetOrCreateUser(beliked).Likes[vid] = true
	UserArray.GetOrCreateUser(beliked).Fans[vid] = true
	UserArray.GetOrCreateUser(vid).Fans[beliked] = true
	UserArray.GetOrCreateUser(vid).Lock_likes.Unlock()
	UserArray.GetOrCreateUser(beliked).Lock_fans.Unlock()
}

// 取消互粉
func (UserArray RelateGraph) Disfriend(vid, beliked uint) {
	UserArray.GetOrCreateUser(vid).Lock_likes.Lock()
	UserArray.GetOrCreateUser(beliked).Lock_fans.Lock()

	delete(UserArray.GetOrCreateUser(vid).Likes, beliked)
	delete(UserArray.GetOrCreateUser(beliked).Likes, vid)
	delete(UserArray.GetOrCreateUser(vid).Fans, beliked)
	delete(UserArray.GetOrCreateUser(beliked).Fans, vid)

	UserArray.GetOrCreateUser(vid).Lock_likes.Unlock()
	UserArray.GetOrCreateUser(beliked).Lock_fans.Unlock()
}

// 2人的关系
// 0:没有关系
// 1：user1 关注了 user2
// 2：user2 关注了 user1
// 3：互粉的好友
func (UserArray RelateGraph) GetRelate(vid1, vid2 uint) int {
	relate := 0
	_, has := UserArray.GetUser(vid1).Likes[vid2]
	if has {
		relate += 1
	}
	_, has2 := UserArray.GetUser(vid1).Fans[vid2]
	if has2 {
		relate += 2
	}
	return relate
}
func (UserArray *RelateGraph) SetRelate(vid1, vid2 uint, relate int) {
	switch relate {
	case 0:
		UserArray.Disfriend(vid1, vid2)
	case 1:
		UserArray.Like(vid1, vid2)
		UserArray.DisLike(vid2, vid1)
	case 2:
		UserArray.Like(vid2, vid1)
		UserArray.DisLike(vid1, vid2)
	case 3:
		UserArray.Makefriend(vid1, vid2)
	}
}
