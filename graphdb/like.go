package graphdb

// 关注的人
func (this User) Getlikes() []uint {
	result := make([]uint, this.Likes.Size())
	i := 0
	for k := range this.Likes.IterKeys() {
		result[i] = k
		i++
	}
	return result
}
