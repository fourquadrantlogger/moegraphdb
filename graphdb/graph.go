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
	return len(this.Users)
}
func (this RelateGraph) GetLikeCountCount() map[int]int {
	likesmap := make(map[int]int, 0)
	for _, v := range this.Users {
		likesmap[len(v.Likes)]++
	}
	return likesmap
}
func (this RelateGraph) GetFanCountCount() map[int]int {
	likesmap := make(map[int]int, 0)
	for _, v := range this.Users {
		likesmap[len(v.Fans)]++
	}
	return likesmap
}
