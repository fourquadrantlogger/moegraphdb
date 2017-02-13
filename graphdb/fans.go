package graphdb

// 粉丝
func (this User) Getfans() []uint {
	result := make([]uint, this.Fans.Size())
	i := 0
	for k := range this.Fans.IterKeys() {
		result[i] = k
		i++
	}
	return result
}
