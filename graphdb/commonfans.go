package graphdb

func (this RelateGraph)GetCommonFans(vid1,vid2 uint)([]uint){
	commonfans:=make([]uint,0)
	for _,f:= range this.Users[vid1].Fans{
		_,hav:=  this.Users[vid1].Fans[f.Uid]
		if(hav){
			commonfans=append(commonfans,f.Uid)
		}
	}
	return commonfans
}
func (this RelateGraph)GetCommonFansCount(vid1,vid2 uint)(int){
	commonfanscount:=0
	for _,f:= range this.Users[vid1].Fans{
		_,hav:=  this.Users[vid1].Fans[f.Uid]
		if(hav){
			commonfanscount++
		}
	}
	return commonfanscount
}
