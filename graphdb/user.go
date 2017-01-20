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

