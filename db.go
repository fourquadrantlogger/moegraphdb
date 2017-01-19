package main

import "sync"

type (
	V struct{
		sync.RWMutex
		vid uint64
		fans map[uint64]*V
		like map[uint64]*V
	}
)

var (
	Vlist []*V
)

func init()  {
	Vlist=make([]*V,3000000000)
}