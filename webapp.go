package main

import (
	"net/http"
	"net/url"
	"fmt"
	"strconv"

	"github.com/timeloveboy/moegraphdb/graphdb"
)

var (
	UserArray graphdb.RelateGraph=graphdb.NewDB(50000000)
)

func main() {
	http.HandleFunc("/like", func(w http.ResponseWriter,r *http.Request){
		m,_:=url.ParseQuery(r.URL.RawQuery)
		vid,_:=strconv.Atoi(m["vid"][0])
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte(fmt.Sprint(UserArray[vid].Getlikes())))
		case http.MethodPost:
			beliked,_:=strconv.Atoi(m["beliked"][0])
			UserArray.Like(uint(vid),uint(beliked))
		case http.MethodDelete:
			beliked,_:=strconv.Atoi(m["beliked"][0])
			UserArray.DisLike(uint(vid),uint(beliked))
		}
	})

	http.HandleFunc("/fans", func(w http.ResponseWriter,r *http.Request){
		m,_:=url.ParseQuery(r.URL.RawQuery)
		vid,_:=strconv.Atoi(m["vid"][0])
		switch r.Method {
		case http.MethodGet:{
			w.Write([]byte(fmt.Sprint(UserArray[vid].Getfans())))
		}
		case http.MethodPost:
			fan,_:=strconv.Atoi(m["fan"][0])
			UserArray.Like(uint(fan),uint(vid))
		case http.MethodDelete:
			fan,_:=strconv.Atoi(m["fan"][0])
			UserArray.DisLike(uint(fan),uint(vid))
		}

	})

	http.HandleFunc("/relate", func(w http.ResponseWriter,r *http.Request){
		m,_:=url.ParseQuery(r.URL.RawQuery)
		vid,_:=strconv.Atoi(m["vid"][0])
		switch r.Method {
		case http.MethodGet:{
			w.Write([]byte(fmt.Sprint(UserArray[vid].Getfans())))
		}
		case http.MethodPost:
			fan,_:=strconv.Atoi(m["fan"][0])
			UserArray[vid].Fans[uint(fan)]=UserArray[uint(fan)]
			UserArray[uint(fan)].Likes[uint(vid)]=UserArray[uint(vid)]
		case http.MethodDelete:
			fan,_:=strconv.Atoi(m["fan"][0])
			delete(UserArray[vid].Fans,uint(fan))
			delete(UserArray[uint(fan)].Likes,uint(vid))
		}

	})

	http.HandleFunc("/relete", func(w http.ResponseWriter,r *http.Request){
		m,_:=url.ParseQuery(r.URL.RawQuery)
		vid1,_:=strconv.Atoi(m["vid1"][0])
		//vid2,_:=strconv.Atoi(m["vid2"][0])
		switch r.Method {
		case http.MethodGet:{
			w.Write([]byte(fmt.Sprint(UserArray[vid1].Getfans())))
		}
		case http.MethodPost:
			fan,_:=strconv.Atoi(m["fan"][0])
			UserArray[vid1].Fans[uint(fan)]=UserArray[uint(fan)]
			UserArray[uint(fan)].Likes[uint(vid1)]=UserArray[uint(vid1)]
		case http.MethodDelete:
			fan,_:=strconv.Atoi(m["fan"][0])
			delete(UserArray[vid1].Fans,uint(fan))
			delete(UserArray[uint(fan)].Likes,uint(vid1))
		}

	})

	fmt.Println("start http server")
	err :=http.ListenAndServe(":8010",nil)
	panic(err)
}
