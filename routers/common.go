package routers

import (
	"encoding/json"
	"fmt"
	"github.com/timeloveboy/moegraphdb/graphdb"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func Handle_common_2_like(w http.ResponseWriter, r *http.Request) {
	moeprint(r)
	bs, _ := ioutil.ReadAll(r.Body)
	user_user := []uint{}
	json.Unmarshal(bs, &user_user)
	if len(user_user) != 2 {
		return
	}
	switch r.Method {
	case http.MethodOptions:
		w.Write([]byte(fmt.Sprint(UserArray.GetCommonLikes(user_user[0], user_user[1]))))
	}
}

func Handle_common_2_fans(w http.ResponseWriter, r *http.Request) {
	moeprint(r)
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	user_user := []uint{}
	json.Unmarshal(bs, &user_user)
	if len(user_user) != 2 {
		return
	}
	switch r.Method {
	case http.MethodOptions:
		w.Write([]byte(fmt.Sprint(UserArray.GetCommonFans(user_user[0], user_user[1]))))
	}
}

func Handle_common_n_like(w http.ResponseWriter, r *http.Request) {
	moeprint(r)
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	users := []uint{}
	json.Unmarshal(bs, &users)

	m, _ := url.ParseQuery(r.URL.RawQuery)
	var min int
	_, have := m["min"]
	if have {
		min, _ = strconv.Atoi(m["min"][0])
	}

	switch r.Method {
	case http.MethodOptions:
		bs, _ := json.Marshal(graphdb.Filtercount_min(UserArray.GetThemCommonLikes(users...), min, 1<<48))
		w.Write(bs)
	}
}

func Handle_common_n_fans(w http.ResponseWriter, r *http.Request) {
	moeprint(r)
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	users := []uint{}
	json.Unmarshal(bs, &users)

	m, _ := url.ParseQuery(r.URL.RawQuery)
	var min int
	_, have := m["min"]
	if have {
		min, _ = strconv.Atoi(m["min"][0])
	}
	switch r.Method {
	case http.MethodOptions:
		bs, _ := json.Marshal(graphdb.Filtercount_min(UserArray.GetThemCommonFans(users...), min, 1<<48))
		w.Write(bs)
	}
}
