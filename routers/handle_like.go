package routers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Handle_like(w http.ResponseWriter, r *http.Request) {
	moeprint(r)
	m, _ := url.ParseQuery(r.URL.RawQuery)
	_, have := m["vid"]
	if !have {
		return
	}
	vid, _ := strconv.Atoi(m["vid"][0])
	switch r.Method {
	case http.MethodGet:
		bs, _ := json.Marshal(UserArray.GetUser(uint(vid)).Getlikes())
		w.Write(bs)
	case http.MethodPost:
		_, have := m["beliked"]
		if !have {
			return
		}
		beliked, _ := strconv.Atoi(m["beliked"][0])
		UserArray.Like(uint(vid), uint(beliked))
	case http.MethodDelete:
		_, have := m["beliked"]
		if !have {
			return
		}
		beliked, _ := strconv.Atoi(m["beliked"][0])
		UserArray.DisLike(uint(vid), uint(beliked))
	}
}
func Handle_like_n(w http.ResponseWriter, r *http.Request) {
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
		i := 0
		for _, v := range relates {
			i++
			UserArray.Like(v.Vid1, v.Vid2)
		}
		w.Write([]byte("import " + strconv.Itoa(i) + " like"))
	}
}
