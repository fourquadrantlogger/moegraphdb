package routers

import (
	"net/http"
	"net/url"
	"strconv"
)

func Handle_user(w http.ResponseWriter, r *http.Request) {
	moeprint(r)
	m, _ := url.ParseQuery(r.URL.RawQuery)
	_, have := m["vid"]
	if !have {
		return
	}
	vid, _ := strconv.Atoi(m["vid"][0])
	switch r.Method {
	case http.MethodGet:
		{
			u := UserArray.GetUser(uint(vid))
			if u != nil {
				w.Write([]byte(UserArray.GetUser(uint(vid)).String()))
			} else {
				w.Write([]byte("no user " + strconv.Itoa(vid)))
			}

		}
	}

}
