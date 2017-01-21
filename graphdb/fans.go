package graphdb

// 粉丝
func (this User) Getfans() []uint {
	result := make([]uint, len(this.Fans))
	i := 0
	for k, _ := range this.Fans {
		result[i] = k
		i++
	}
	return result
}
