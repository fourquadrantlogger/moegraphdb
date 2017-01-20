package graphdb


// 关注他
func (UserArray RelateGraph)Like(vid,beliked uint){
	UserArray[int(vid)].Likes[beliked]=UserArray[beliked]
	UserArray[int(beliked)].Fans[vid]=UserArray[vid]
}
// 取消关注他
func (UserArray RelateGraph)DisLike(vid,beliked uint){
	delete(UserArray[int(vid)].Likes,beliked)
	delete(UserArray[int(beliked)].Fans,vid)
}
// 互粉
func (UserArray RelateGraph)Makefriend(vid,beliked uint){
	UserArray[int(vid)].Likes[beliked]=UserArray[beliked]
	UserArray[int(beliked)].Likes[vid]=UserArray[vid]
	UserArray[int(beliked)].Fans[vid]=UserArray[vid]
	UserArray[int(vid)].Fans[beliked]=UserArray[beliked]
}
// 取消互粉
func (UserArray RelateGraph)Disfriend(vid,beliked uint){
	delete(UserArray[int(vid)].Likes,beliked)
	delete(UserArray[int(beliked)].Likes,vid)
	delete(UserArray[int(vid)].Fans,beliked)
	delete(UserArray[int(beliked)].Fans,vid)
}

// 2人的关系
// 0:没有关系
// 1：user1 关注了 user2
// 2：user2 关注了 user1
// 3：互粉的好友
func (UserArray RelateGraph)GetRelate(vid1,vid2 uint)int{
	relate:=0
	_,has:=UserArray[vid1].Likes[uint(vid2)]
	if(has){
		relate+=1
	}
	_,has2:=UserArray[vid1].Fans[uint(vid2)]
	if(has2){
		relate+=2
	}
	return relate
}
func (UserArray *RelateGraph)SetRelate(vid1,vid2 uint,relate int){
	switch relate {
	case 0:UserArray.Disfriend(vid1,vid2)
	case 1:UserArray.Like(vid1,vid2)
	case 2:UserArray.Like(vid2,vid1)
	case 3:UserArray.Makefriend(vid1,vid2)
	}
}