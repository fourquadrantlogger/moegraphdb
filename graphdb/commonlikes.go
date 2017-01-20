package graphdb
func (this RelateGraph)GetCommonLikes(vid1,vid2 uint)([]uint){
	commonlikes:=make([]uint,0)
	for _,f:= range this.Users[vid1].Likes{
		_,hav:=  this.Users[vid1].Likes[f.Uid]
		if(hav){
			commonlikes=append(commonlikes,f.Uid)
		}
	}
	return commonlikes
}
func (this RelateGraph)GetCommonLikesCount(vid1,vid2 uint)(int){
	commonlikescount:=0
	for _,f:= range this.Users[vid1].Likes{
		_,hav:=  this.Users[vid1].Likes[f.Uid]
		if(hav){
			commonlikescount++
		}
	}
	return commonlikescount
}
