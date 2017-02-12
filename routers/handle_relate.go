package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Handle_relate(w http.ResponseWriter, r *http.Request) {
	moeprint(r)
	m, _ := url.ParseQuery(r.URL.RawQuery)
	_, have := m["vid1"]
	if !have {
		return
	}
	_, have2 := m["vid2"]
	if !have2 {
		return
	}
	vid1, _ := strconv.Atoi(m["vid1"][0])
	vid2, _ := strconv.Atoi(m["vid2"][0])
	switch r.Method {
	case http.MethodGet:
		{
			w.Write([]byte(fmt.Sprint(UserArray.GetRelate(uint(vid1), uint(vid2)))))
		}
	case http.MethodPost:
		relate, _ := strconv.Atoi(m["relate"][0])
		UserArray.SetRelate(uint(vid1), uint(vid2), relate)
		w.Write([]byte("ok"))
	case http.MethodDelete:
		UserArray.Disfriend(uint(vid1), uint(vid2))
		w.Write([]byte("ok"))
	}
}

func Handle_relate_n(w http.ResponseWriter, r *http.Request) {
	moeprint(r)
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	m, _ := url.ParseQuery(r.URL.RawQuery)
	_, have := m["type"]
	if !have {
		return
	}

	relates := make([]User_Fans_Relate, 0)
	switch m["type"][0] {
	case "json":
		json.Unmarshal(bs, &relates)
	case "row":
		lines := strings.Split(string(bs), "\n")
		for i, _ := range lines {
			user_fan := strings.Split(lines[i], ",")
			vid1, _ := strconv.Atoi(user_fan[0])
			vid2, _ := strconv.Atoi(user_fan[1])
			relate, _ := strconv.Atoi(user_fan[2])
			relates = append(relates, User_Fans_Relate{Vid1: uint(vid1), Vid2: uint(vid2), Relate: relate})
		}
	}

	switch r.Method {
	case http.MethodPost:
		for _, v := range relates {
			UserArray.SetRelate(v.Vid1, v.Vid2, v.Relate)
		}
		w.Write([]byte("ok"))
	}
}
