package graphdb

// 找到2个用户，共同的粉丝
func (this RelateGraph) GetCommonFans(vid1, vid2 uint) []uint {
	commonfans := make([]uint, 0)
	for f := range this.GetUser(vid1).Fans.IterKeys() {
		hav := this.GetUser(vid2).Fans.Has(f)
		if hav {
			commonfans = append(commonfans, f)
		}
	}

	return commonfans
}

// 找到n个用户的粉丝，人/数
func (this RelateGraph) GetThemCommonFans(vids ...uint) map[uint]int {
	likesmap := make(map[uint]int, 0)
	for _, v := range vids {
		for f := range this.GetUser(v).Fans.IterKeys() {
			likesmap[f]++
		}
	}
	return likesmap
}
