package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"encoding/json"
	"github.com/timeloveboy/moegraphdb/graphdb"
	"io/ioutil"
	"strings"
)

type (
	User_Fans struct {
		Vid1 uint `json:"vid1"`
		Vid2 uint `json:"vid2"`
	}
	User_Fans_Relate struct {
		Vid1   uint `json:"vid1"`
		Vid2   uint `json:"vid2"`
		Relate int  `json:"relate"`
	}
)

var (
	UserArray graphdb.RelateGraph
	OpenLog   bool = true
)

func moeprint(r *http.Request) {
	if OpenLog {
		fmt.Println(r.Method, "\t", r.URL)
	}
}
func main() {
	UserArray = graphdb.NewDB()
	http.HandleFunc("/like", func(w http.ResponseWriter, r *http.Request) {
		moeprint(r)
		m, _ := url.ParseQuery(r.URL.RawQuery)
		_, have := m["vid"]
		if !have {
			return
		}
		vid, _ := strconv.Atoi(m["vid"][0])
		switch r.Method {
		case http.MethodGet:
			bs, _ := json.Marshal(UserArray.Users[uint(vid)].Getlikes())
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
	})

	http.HandleFunc("/like/n", func(w http.ResponseWriter, r *http.Request) {
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
				UserArray.Like(v.Vid1, v.Vid2)
			}
			w.Write([]byte("ok"))
		}
	})
	http.HandleFunc("/fans", func(w http.ResponseWriter, r *http.Request) {
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
				bs, _ := json.Marshal(UserArray.GetUser(uint(vid)).Getfans())
				w.Write(bs)
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

	})

	http.HandleFunc("/fans/n", func(w http.ResponseWriter, r *http.Request) {
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
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
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
				w.Write([]byte(fmt.Sprint(UserArray.GetUser(uint(vid)))))
			}
		}

	})

	http.HandleFunc("/relate", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.HandleFunc("/relate/n", func(w http.ResponseWriter, r *http.Request) {
		moeprint(r)
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
	})
	http.HandleFunc("/common/2/likes", func(w http.ResponseWriter, r *http.Request) {
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
	})
	http.HandleFunc("/common/2/fans", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.HandleFunc("/common/n/likes", func(w http.ResponseWriter, r *http.Request) {
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
			bs, _ := json.Marshal(graphdb.Filtercount_min(UserArray.GetThemCommonLikes(users...), min))
			w.Write(bs)
		}
	})

	http.HandleFunc("/filter/n", func(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println(users)
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
	})
	http.HandleFunc("/common/n/fans", func(w http.ResponseWriter, r *http.Request) {
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
			bs, _ := json.Marshal(graphdb.Filtercount_min(UserArray.GetThemCommonFans(users...), min))
			w.Write(bs)
		}
	})

	http.HandleFunc("/count/relate", func(w http.ResponseWriter, r *http.Request) {
		moeprint(r)
		switch r.Method {
		case http.MethodGet:
			bs, _ := json.Marshal(UserArray.GetUserRelateCount())
			w.Write(bs)
		}
	})

	http.HandleFunc("/count/user", func(w http.ResponseWriter, r *http.Request) {
		moeprint(r)

		switch r.Method {
		case http.MethodGet:
			bs, _ := json.Marshal(UserArray.GetUserUserCount())
			w.Write(bs)
		}
	})

	http.HandleFunc("/count/like", func(w http.ResponseWriter, r *http.Request) {
		moeprint(r)

		switch r.Method {
		case http.MethodGet:
			bs, _ := json.Marshal(UserArray.GetLikeCountCount())
			w.Write(bs)
		}
	})
	http.HandleFunc("/count/fans", func(w http.ResponseWriter, r *http.Request) {
		moeprint(r)
		switch r.Method {
		case http.MethodGet:
			bs, _ := json.Marshal(UserArray.GetFanCountCount())
			w.Write(bs)
		}
	})
	fmt.Println("start http server")
	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		panic(err)
	}
}
