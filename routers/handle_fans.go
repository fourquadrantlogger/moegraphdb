package routers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Handle_fans(w http.ResponseWriter, r *http.Request) {
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
				bs, _ := json.Marshal(u.Getfans())
				w.Write(bs)
			} else {
				w.Write([]byte("no user " + strconv.Itoa(vid)))
			}

		}
	case http.MethodPost:
		_, have := m["fan"]
		if !have {
			return
		}
		fan, _ := strconv.Atoi(m["fan"][0])
		UserArray.Like(uint(fan), uint(vid))
	case http.MethodDelete:
		_, have := m["fan"]
		if !have {
			return
		}
		fan, _ := strconv.Atoi(m["fan"][0])
		UserArray.DisLike(uint(fan), uint(vid))
	}

}

func Handle_fans_n(w http.ResponseWriter, r *http.Request) {
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
	relates := make([]User_Fans, 0)
	switch m["type"][0] {
	case "json":
		json.Unmarshal(bs, &relates)
	case "row":
		lines := strings.Split(string(bs), "\n")
		for i, _ := range lines {
			user_fan := strings.Split(lines[i], ",")
			vid1, _ := strconv.Atoi(user_fan[0])
			vid2, _ := strconv.Atoi(user_fan[1])
			relates = append(relates, User_Fans{Vid1: uint(vid1), Vid2: uint(vid2)})
		}
	}

	switch r.Method {
	case http.MethodPost:
		for _, v := range relates {
			UserArray.Like(v.Vid2, v.Vid1)
		}
		w.Write([]byte("ok"))
	}
}
