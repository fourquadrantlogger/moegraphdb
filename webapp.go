package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"encoding/json"
	"github.com/timeloveboy/moegraphdb/graphdb"
	"io/ioutil"
	"os"
)

var (
	graphusercount = os.Getenv("usercount")
	UserArray      graphdb.RelateGraph
)

func main() {
	cap, err := strconv.Atoi(graphusercount)
	if err != nil {
		panic(err)
	}
	UserArray = graphdb.NewDB(cap)
	http.HandleFunc("/like", func(w http.ResponseWriter, r *http.Request) {
		m, _ := url.ParseQuery(r.URL.RawQuery)
		_, have := m["vid"]
		if !have {
			return
		}
		vid, _ := strconv.Atoi(m["vid"][0])
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte(fmt.Sprint(UserArray.Users[uint(vid)].Getlikes())))
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

	http.HandleFunc("/fans", func(w http.ResponseWriter, r *http.Request) {
		m, _ := url.ParseQuery(r.URL.RawQuery)
		_, have := m["vid"]
		if !have {
			return
		}
		vid, _ := strconv.Atoi(m["vid"][0])
		switch r.Method {
		case http.MethodGet:
			{
				w.Write([]byte(fmt.Sprint(UserArray.GetUser(uint(vid)).Getfans())))
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

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
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
			//case http.MethodPost:
			//	body, _ := ioutil.ReadAll(r.Body)
			//	var info map[string]interface{}
			//	err := json.Unmarshal(body, &info)
			//	panic(err)
			//	UserArray.GetUser(uint(vid)).Info = info
			//case http.MethodPut:
			//	body, _ := ioutil.ReadAll(r.Body)
			//	var info map[string]interface{}
			//	err := json.Unmarshal(body, &info)
			//	panic(err)
			//	for k, v := range info {
			//		UserArray.GetUser(uint(vid)).Info[k] = v
			//	}
		}

	})

	http.HandleFunc("/relate/2", func(w http.ResponseWriter, r *http.Request) {
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
		case http.MethodDelete:
			UserArray.Disfriend(uint(vid1), uint(vid2))
		}
	})

	http.HandleFunc("/relate/n", func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		relates := []struct {
			Vid1   uint `json:"vid1"`
			Vid2   uint `json:"vid2"`
			Relate int  `json:"relate"`
		}{}
		json.Unmarshal(bs, &relates)
		switch r.Method {
		case http.MethodPost:
			for _, v := range relates {
				UserArray.SetRelate(v.Vid1, v.Vid2, v.Relate)
			}
		}
	})
	http.HandleFunc("/common/2/likes", func(w http.ResponseWriter, r *http.Request) {
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

			bs, _ := json.Marshal(graphdb.Filter_min(UserArray.GetThemCommonLikes(users...), min))
			w.Write(bs)
		}
	})
	http.HandleFunc("/common/n/fans", func(w http.ResponseWriter, r *http.Request) {
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
			bs, _ := json.Marshal(graphdb.Filter_min(UserArray.GetThemCommonFans(users...), min))
			w.Write(bs)
		}
	})

	http.HandleFunc("/relate/count", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			bs, _ := json.Marshal(UserArray.GetUserRelateCount())
			w.Write(bs)
		}
	})
	fmt.Println("start http server")
	err = http.ListenAndServe(":8010", nil)
	if err != nil {
		panic(err)
	}
}
