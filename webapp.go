package main

import (
	"net/http"
	"net/url"
	"fmt"
	"strconv"

	"github.com/timeloveboy/moegraphdb/graphdb"
	"io/ioutil"
	"encoding/json"
	"os"
)

var (
	graphusercount=os.Getenv("usercount")
	UserArray graphdb.RelateGraph
)

func main() {
	cap,err:=strconv.Atoi(graphusercount)
	if(err!=nil) {
		panic(err)
	}
	UserArray=graphdb.NewDB(cap)
	http.HandleFunc("/like", func(w http.ResponseWriter,r *http.Request){
		m,_:=url.ParseQuery(r.URL.RawQuery)
		_,have:=m["vid"]
		if(!have){
			return
		}
		vid,_:=strconv.Atoi(m["vid"][0])
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte(fmt.Sprint(UserArray.Users[vid].Getlikes())))
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
			w.Write([]byte(fmt.Sprint(UserArray.Users[vid].Getfans())))
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
			w.Write([]byte(fmt.Sprint(UserArray.Users[vid])))
		}
		case http.MethodPost:
			body,_:=ioutil.ReadAll(r.Body)
			var info map[string]interface{}
			err:=json.Unmarshal(body,&info)
			panic(err)
			UserArray.Users[uint(vid)].Info=info
		case http.MethodPut:
			body,_:=ioutil.ReadAll(r.Body)
			var info map[string]interface{}
			err:=json.Unmarshal(body,&info)
			panic(err)
			for k,v:=range info{
				UserArray.Users[uint(vid)].Info[k]=v
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

	http.HandleFunc("/common/2/likes", func(w http.ResponseWriter,r *http.Request){
		bs,_:=ioutil.ReadAll(r.Body)
		user_user:=[]uint{}
		json.Unmarshal(bs,&user_user)
		if(len(user_user)!=2){
			return
		}
		switch r.Method {
		case http.MethodOptions:
			w.Write([]byte(fmt.Sprint(UserArray.GetCommonLikes(user_user[0],user_user[1]))))
		}
	})
	http.HandleFunc("/common/2/fans", func(w http.ResponseWriter,r *http.Request){
		bs,_:=ioutil.ReadAll(r.Body)
		user_user:=[]uint{}
		json.Unmarshal(bs,&user_user)
		if(len(user_user)!=2){
			return
		}
		switch r.Method {
		case http.MethodOptions:
			w.Write([]byte(fmt.Sprint(UserArray.GetCommonFans(user_user[0],user_user[1]))))
		}
	})

	http.HandleFunc("/common/n/likes", func(w http.ResponseWriter,r *http.Request){
		bs,_:=ioutil.ReadAll(r.Body)
		users:=[]uint{}
		json.Unmarshal(bs,&users)
		switch r.Method {
		case http.MethodOptions:
			bs,_:=json.Marshal(UserArray.GetThemCommonLikes(users...))
			w.Write(bs)
		}
	})
	http.HandleFunc("/common/n/fans", func(w http.ResponseWriter,r *http.Request){
		bs,_:=ioutil.ReadAll(r.Body)
		users:=[]uint{}
		json.Unmarshal(bs,&users)
		switch r.Method {
		case http.MethodOptions:
			bs,_:=json.Marshal(UserArray.GetThemCommonFans(users...))
			w.Write(bs)
		}
	})
	fmt.Println("start http server")
	err=http.ListenAndServe(":8010",nil)
	if(err!=nil) {
		panic(err)
	}
}
