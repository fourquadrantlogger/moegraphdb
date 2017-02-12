package graphdb

func Filtercount_min(source map[uint]int, min, max int) map[uint]int {
	for k, v := range source {
		if v < min {
			delete(source, k)
		}
		if v > max {
			delete(source, k)
		}
	}
	return source
}
