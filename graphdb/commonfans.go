package graphdb

func (this RelateGraph)GetCommonFans(vid1,vid2 uint)([]uint){
	commonfans:=make([]uint,0)
	for _,f:= range this[vid1].Fans{
		_,hav:=  this[vid1].Fans[f.Uid]
		if(hav){
			commonfans=append(commonfans,f.Uid)
		}
	}
	return commonfans
}
func (this RelateGraph)GetCommonFansCount(vid1,vid2 uint)(int){
	commonfanscount:=0
	for _,f:= range this[vid1].Fans{
		_,hav:=  this[vid1].Fans[f.Uid]
		if(hav){
			commonfanscount++
		}
	}
	return commonfanscount
}
func (this RelateGraph)GetCommonLikes(vid1,vid2 uint)([]uint){
	commonlikes:=make([]uint,0)
	for _,f:= range this[vid1].Likes{
		_,hav:=  this[vid1].Likes[f.Uid]
		if(hav){
			commonlikes=append(commonlikes,f.Uid)
		}
	}
	return commonlikes
}
func (this RelateGraph)GetCommonLikesCount(vid1,vid2 uint)(int){
	commonlikescount:=0
	for _,f:= range this[vid1].Likes{
		_,hav:=  this[vid1].Likes[f.Uid]
		if(hav){
			commonlikescount++
		}
	}
	return commonlikescount
}
