package graphdb


type RelateGraph []*User
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
