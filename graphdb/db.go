package graphdb

import "fmt"

func NewDB(count int) RelateGraph {
	UserArray := make(map[uint]*User, count)
	fmt.Println("make map user:", count)
	return RelateGraph{
		Users:  UserArray,
		Indexs: make(map[string]map[string]interface{}, 0),
	}
}
