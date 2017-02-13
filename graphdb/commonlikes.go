package graphdb

// 找到2个用户，共同关注的人
func (this RelateGraph) GetCommonLikes(vid1, vid2 uint) []uint {
	commonlikes := make([]uint, 0)
	for f := range this.GetUser(vid1).Likes.IterKeys() {
		hav := this.GetUser(vid2).Likes.Has(f)
		if hav {
			commonlikes = append(commonlikes, f)
		}
	}
	return commonlikes
}

// 找到n个用户，关注的人/数
func (this RelateGraph) GetThemCommonLikes(vids ...uint) map[uint]int {
	likesmap := make(map[uint]int, 0)
	for _, v := range vids {
		for k := range this.GetUser(v).Likes.IterKeys() {
			likesmap[k]++
		}
	}
	return likesmap
}
