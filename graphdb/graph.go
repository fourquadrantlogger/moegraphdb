package graphdb


type RelateGraph []*User

func (this RelateGraph)GetUserRelateCount()(int){
	relatecount:=0
	for _,v:= range this{
		relatecount+=len(v.Likes)
	}
	return relatecount
}

func (this RelateGraph)SearchUserWhereInfo(key string,value interface{}){

}