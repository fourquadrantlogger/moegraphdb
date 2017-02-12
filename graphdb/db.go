package graphdb

func NewDB() RelateGraph {
	UserArray := syncmapNew()
	return RelateGraph{
		users:  *UserArray,
		Indexs: make(map[string]map[string]interface{}, 0),
	}
}
