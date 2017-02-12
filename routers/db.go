package routers

import "github.com/timeloveboy/moegraphdb/graphdb"

type (
	User_Fans struct {
		Vid1 uint `json:"vid1"`
		Vid2 uint `json:"vid2"`
	}
	User_Fans_Relate struct {
		Vid1   uint `json:"vid1"`
		Vid2   uint `json:"vid2"`
		Relate int  `json:"relate"`
	}
)

var (
	UserArray graphdb.RelateGraph
)
