package graphdb

func Filter_min(source map[uint]int, min int) map[uint]int {
	for k, v := range source {
		if v < min {
			delete(source, k)
		}
	}
	return source
}
