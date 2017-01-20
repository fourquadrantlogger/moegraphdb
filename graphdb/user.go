package graphdb

import "sync"

type (
	User struct{
		Lock_fans  sync.RWMutex
		Lock_likes sync.RWMutex
		Uid        uint
		Fans       map[uint]*User
		Likes      map[uint]*User
	}
)
// 粉丝数
func (this *User)FansCount()(int){
	if(this!=nil&&this.Fans!=nil){
		return len(this.Fans)
	}
	return 0
}
// 关注数
func (this *User)LikesCount()(int){
	if(this!=nil&&this.Likes!=nil){
		return len(this.Likes)
	}
	return 0
}
