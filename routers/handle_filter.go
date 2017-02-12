package routers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func Handle_filter(w http.ResponseWriter, r *http.Request) {
	moeprint(r)
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	users := make([]uint, 0)
	json.Unmarshal(bs, &users)

	m, _ := url.ParseQuery(r.URL.RawQuery)
	var filtertype string
	_, have := m["type"]
	if have {
		filtertype = m["type"][0]
	}
	var max, min int
	_, have = m["max"]
	if have {
		max, _ = strconv.Atoi(m["max"][0])
	}
	_, have = m["min"]
	if have {
		min, _ = strconv.Atoi(m["min"][0])
	}

	switch r.Method {
	case http.MethodOptions:
		bs := make([]byte, 0)
		switch filtertype {
		case "fanscount":
			bs, _ = json.Marshal(UserArray.Filterusers_fanscount(users, max, min))
		case "likescount":
			bs, _ = json.Marshal(UserArray.Filterusers_likescount(users, max, min))
		}
		w.Write(bs)
	}
}
