package graphdb

// 找到2个用户，共同关注的人
func (this RelateGraph) GetCommonLikes(vid1, vid2 uint) []uint {
	commonlikes := make([]uint, 0)
	for _, f := range this.Users[vid1].Likes {
		_, hav := this.Users[vid1].Likes[f.Uid]
		if hav {
			commonlikes = append(commonlikes, f.Uid)
		}
	}
	return commonlikes
}

// 找到n个用户，关注的人/数
func (this RelateGraph) GetThemCommonLikes(vids ...uint) map[uint]int {
	likesmap := make(map[uint]int, 0)
	for _, v := range vids {
		for k, _ := range this.Users[v].Likes {
			likesmap[k]++
		}
	}
	return likesmap
}
