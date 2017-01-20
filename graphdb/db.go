package graphdb


func NewDB(count uint)(RelateGraph){
	UserArray:=make([]*User,count)
	for i,_:=range UserArray{
		UserArray[i]=new(User)
		UserArray[i].Uid =uint(i)
		UserArray[i].Fans=make( map[uint]*User )
		UserArray[i].Likes=make( map[uint]*User )
	}
	return UserArray
}

