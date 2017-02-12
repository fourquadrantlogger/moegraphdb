package graphdb

type RelateGraph struct {
	//用户
	users SyncMap
	//info 的索引
	Indexs map[string](map[string]interface{})
}

func (this *RelateGraph) InitIndex() {

}
func (this RelateGraph) GetUserRelateCount() int {
	relatecount := 0
	for v := range this.users.IterItems() {
		relatecount += len(v.Value.Likes)
	}
	return relatecount
}
func (this RelateGraph) GetUserUserCount() int {
	return this.users.Size()
}
func (this RelateGraph) GetLikeCountCount() map[int]int {
	likesmap := make(map[int]int, 0)
	for v := range this.users.IterItems() {
		likesmap[len(v.Value.Likes)]++
	}

	return likesmap
}
func (this RelateGraph) GetFanCountCount() map[int]int {
	likesmap := make(map[int]int, 0)
	for v := range this.users.IterItems() {
		likesmap[len(v.Value.Fans)]++
	}
	return likesmap
}
