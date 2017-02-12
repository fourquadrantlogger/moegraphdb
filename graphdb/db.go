package graphdb

func NewDB() RelateGraph {
	UserArray := SyncmapNew()
	return RelateGraph{
		Users:  *UserArray,
		Indexs: make(map[string]map[string]interface{}, 0),
	}
}
