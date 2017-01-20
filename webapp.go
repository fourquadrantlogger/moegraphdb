package main

import (
	"net/http"
	"net/url"
	"fmt"
	"strconv"

	"github.com/timeloveboy/moegraphdb/graphdb"
	"io/ioutil"
	"encoding/json"
)

var (
	UserArray graphdb.RelateGraph=graphdb.NewDB(50000000)
)

func main() {
	http.HandleFunc("/like", func(w http.ResponseWriter,r *http.Request){
		m,_:=url.ParseQuery(r.URL.RawQuery)
		_,have:=m["vid"]
		if(!have){
			return
		}
		vid,_:=strconv.Atoi(m["vid"][0])
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte(fmt.Sprint(UserArray[vid].Getlikes())))
		case http.MethodPost:
			_,have:=m["beliked"]
			if(!have){
				return
			}
			beliked,_:=strconv.Atoi(m["beliked"][0])
			UserArray.Like(uint(vid),uint(beliked))
		case http.MethodDelete:
			_,have:=m["beliked"]
			if(!have){
				return
			}
			beliked,_:=strconv.Atoi(m["beliked"][0])
			UserArray.DisLike(uint(vid),uint(beliked))
		}
	})

	http.HandleFunc("/fans", func(w http.ResponseWriter,r *http.Request){
		m,_:=url.ParseQuery(r.URL.RawQuery)
		_,have:=m["vid"]
		if(!have){
			return
		}
		vid,_:=strconv.Atoi(m["vid"][0])
		switch r.Method {
		case http.MethodGet:{
			w.Write([]byte(fmt.Sprint(UserArray[vid].Getfans())))
		}
		case http.MethodPost:
			_,have:=m["fan"]
			if(!have){
				return
			}
			fan,_:=strconv.Atoi(m["fan"][0])
			UserArray.Like(uint(fan),uint(vid))
		case http.MethodDelete:
			_,have:=m["fan"]
			if(!have){
				return
			}
			fan,_:=strconv.Atoi(m["fan"][0])
			UserArray.DisLike(uint(fan),uint(vid))
		}

	})

	http.HandleFunc("/user", func(w http.ResponseWriter,r *http.Request){
		m,_:=url.ParseQuery(r.URL.RawQuery)
		_,have:=m["vid"]
		if(!have){
			return
		}
		vid,_:=strconv.Atoi(m["vid"][0])
		switch r.Method {
		case http.MethodGet:{
			w.Write([]byte(fmt.Sprint(UserArray[vid])))
		}
		case http.MethodPost:
			body,_:=ioutil.ReadAll(r.Body)
			var info map[string]interface{}
			err:=json.Unmarshal(body,&info)
			panic(err)
			UserArray[uint(vid)].Info=info
		case http.MethodPut:
			body,_:=ioutil.ReadAll(r.Body)
			var info map[string]interface{}
			err:=json.Unmarshal(body,&info)
			panic(err)
			for k,v:=range info{
				UserArray[uint(vid)].Info[k]=v
			}
		}

	})

	http.HandleFunc("/relate", func(w http.ResponseWriter,r *http.Request){
		m,_:=url.ParseQuery(r.URL.RawQuery)
		_,have:=m["vid1"]
		if(!have){
			return
		}
		_,have2:=m["vid2"]
		if(!have2){
			return
		}
		vid1,_:=strconv.Atoi(m["vid1"][0])
		vid2,_:=strconv.Atoi(m["vid2"][0])
		switch r.Method {
		case http.MethodGet:{
			w.Write([]byte(fmt.Sprint(UserArray.GetRelate(uint(vid1),uint(vid2))) ))
		}
		case http.MethodPost:
			relate,_:=strconv.Atoi(m["relate"][0])
			UserArray.SetRelate(uint(vid1),uint(vid2),relate)
		case http.MethodDelete:
			UserArray.Disfriend(uint(vid1),uint(vid2))
		}

	})

	fmt.Println("start http server")
	err :=http.ListenAndServe(":8010",nil)
	panic(err)
}
