package graphdb

// 关注的人
func (this User) Getlikes() []uint {
	result := make([]uint, len(this.Likes))
	i := 0
	for k, _ := range this.Likes {
		result[i] = k
		i++
	}
	return result
}
