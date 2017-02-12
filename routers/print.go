package routers

import (
	"fmt"
	"net/http"
)

var (
	OpenLog bool = true
)

func moeprint(r *http.Request) {
	if OpenLog {
		fmt.Println(r.Method, "\t", r.URL)
	}
}
