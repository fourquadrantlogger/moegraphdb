package graphdb

type RelateGraph struct {
	//用户
	Users map[uint]*User
	//info 的索引
	Indexs map[string](map[string]interface{})
}

func (this *RelateGraph) InitIndex() {

}
func (this RelateGraph) GetUserRelateCount() int {
	relatecount := 0
	for _, v := range this.Users {
		relatecount += len(v.Likes)
	}
	return relatecount
}
func (this RelateGraph) GetUserUserCount() int {
	relatecount := 0
	for range this.Users {
		relatecount++
	}
	return relatecount
}
func (this RelateGraph) SearchUserWhereInfo(key string, value interface{}) {

}
