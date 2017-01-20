package graphdb


func NewDB(count int)(RelateGraph){
	UserArray:=make([]*User,count)
	for i,_:=range UserArray{
		UserArray[i]=new(User)
		UserArray[i].Uid =uint(i)
		UserArray[i].Fans=make( map[uint]*User,0)
		UserArray[i].Likes=make( map[uint]*User,0)
	}
	return RelateGraph{
		Users:UserArray,
		Indexs:make(map[string]map[string]interface{},0),
	}
}

