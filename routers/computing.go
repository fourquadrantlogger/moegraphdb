package routers

import (
	"encoding/json"
	"github.com/timeloveboy/moegraphdb/computing"
	"github.com/timeloveboy/moegraphdb/graphdb"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func Handle_computing_deadfans(w http.ResponseWriter, r *http.Request) {
	moeprint(r)
	m, _ := url.ParseQuery(r.URL.RawQuery)
	var vid int
	_, have := m["vid"]
	if have {
		vid, _ = strconv.Atoi(m["vid"][0])
	} else {
		w.Write([]byte("require vid"))
	}

	var fansmax int = 1000000
	_, have = m["fansmax"]
	if have {
		fansmax, _ = strconv.Atoi(m["fansmax"][0])
	}

	var existcount int = 10
	_, have = m["existcount"]
	if have {
		existcount, _ = strconv.Atoi(m["existcount"][0])
	}

	switch r.Method {
	case http.MethodGet:
		v := UserArray.GetUser(uint(vid))
		if v == nil {
			w.Write([]byte("no user " + strconv.Itoa(vid)))
			return
		}
		vid_likes := v.Getlikes()
		vid_likes_min1000000 := UserArray.Filterusers_fanscount(vid_likes, fansmax, 0)
		count_count := UserArray.GetThemCommonFans(vid_likes_min1000000...)
		bs, _ := json.Marshal(graphdb.Filtercount_min(count_count, existcount, 1<<32))
		w.Write(bs)
	}
}

func AutoComputing(w http.ResponseWriter, r *http.Request) {
	moeprint(r)
	m, _ := url.ParseQuery(r.URL.RawQuery)
	var fansmax int = 1000000
	_, have := m["fansmax"]
	if have {
		fansmax, _ = strconv.Atoi(m["fansmax"][0])
	}

	var existcount int = 10
	_, have = m["existcount"]
	if have {
		existcount, _ = strconv.Atoi(m["existcount"][0])
	}

	var taskname string = "result"
	_, have = m["taskname"]
	if have {
		taskname = m["existcount"][0]
	}
	ids := make([]int, 0)
	switch r.Method {
	case http.MethodPost:
		bs, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(bs, &ids)

		if computing.Start == false {
			computing.Start = true
			go computing.Mapper(UserArray, fansmax, existcount, ids, taskname)
		}
	}
	w.Write([]byte("now_vid:" + strconv.Itoa(computing.Now_vid)))
}
