package graphdb

import (
	"fmt"
	"strconv"
)

type (
	User struct {
		//Info map[string]interface{}
		Uid   uint
		Fans  *SafeMap
		Likes *SafeMap
	}
)

func (this *RelateGraph) GetUser(vid uint) *User {
	if u, have := this.Users.Get(vid); have {
		return u
	}
	return nil
}
func (this *RelateGraph) GetOrCreateUser(vid uint) *User {
	if u, have := this.Users.Get(vid); have {
		return u
	} else {

		this.CreateUser(vid)
		return this.GetUser(vid)
	}
}

func (this *RelateGraph) CreateUser(vid uint) {
	if _, have := this.Users.Get(vid); have {
		panic("user exist" + strconv.Itoa(int(vid)))
		return
	} else {
		this.Users.Set(vid, &User{Uid: vid,
			Fans:  SafemapNewWithShard(1),
			Likes: SafemapNewWithShard(1),
		})
	}
}

func (this *User) String() string {
	//info, _ := json.Marshal(this.Info)
	//info := ""
	return "{ \"Uid\":" + fmt.Sprint(this.Uid) + ",\"FansCount\":" + fmt.Sprint(this.FansCount()) + ",\"LikesCount\":" + fmt.Sprint(this.LikesCount()) + "}"
}

// 粉丝数
func (this *User) FansCount() int {
	if this != nil && this.Fans != nil {
		return this.Fans.Size()
	}
	return 0
}

// 关注数
func (this *User) LikesCount() int {
	if this != nil && this.Likes != nil {
		return this.Likes.Size()
	}
	return 0
}

// 用户更多信息
//func (this *User) SetInfo(info map[string]interface{}) {
//	this.Info = info
//}
//
//func (this *User) GetInfo() (info map[string]interface{}) {
//	info = this.Info
//	return
//}
