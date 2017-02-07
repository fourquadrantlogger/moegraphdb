package graphdb

func (this *RelateGraph) Filterusers_fanscount(userlist []uint, max, min int) []uint {
	result := []uint{}
	for _, v := range userlist {
		u := this.GetUser(v).FansCount()
		if u < max && u > min {
			result = append(result, v)
		}
	}
	return result
}

func (this *RelateGraph) Filterusers_likescount(userlist []uint, max, min int) []uint {
	result := []uint{}
	for _, v := range userlist {
		u := this.GetUser(v).LikesCount()
		if u < max && u > min {
			result = append(result, v)
		}
	}
	return result
}
