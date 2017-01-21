package graphdb

func NewDB() RelateGraph {
	UserArray := make(map[uint]*User)
	return RelateGraph{
		Users:  UserArray,
		Indexs: make(map[string]map[string]interface{}, 0),
	}
}
