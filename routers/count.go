package routers

import (
	"encoding/json"
	"net/http"
)

func Handle_count_relate(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		bs, _ := json.Marshal(UserArray.GetUserRelateCount())
		w.Write(bs)
	}
}
func Handle_count_user(w http.ResponseWriter, r *http.Request) {
	moeprint(r)

	switch r.Method {
	case http.MethodGet:
		bs, _ := json.Marshal(UserArray.GetUserUserCount())
		w.Write(bs)
	}
}

func Handle_count_like(w http.ResponseWriter, r *http.Request) {
	moeprint(r)

	switch r.Method {
	case http.MethodGet:
		bs, _ := json.Marshal(UserArray.GetLikeCountCount())
		w.Write(bs)
	}
}

func Handle_count_fans(w http.ResponseWriter, r *http.Request) {
	moeprint(r)
	switch r.Method {
	case http.MethodGet:
		bs, _ := json.Marshal(UserArray.GetFanCountCount())
		w.Write(bs)
	}
}
