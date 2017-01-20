package main

import (
	"net/http"
	"net/url"
	"fmt"
	"strconv"

	"sync"
)

type (
	User struct{
		Lock_fans sync.RWMutex
		Lock_likes sync.RWMutex
		Vid uint
		Fans map[uint]*User
		Likes map[uint]*User
	}
)

var (
	UserArray []*User
)

func init()  {
	UserArray =make([]*User,50000000)
	for i,_:=range UserArray{
		UserArray[i]=new(User)
		UserArray[i].Vid=uint(i)
		UserArray[i].Fans=make( map[uint]*User )
		UserArray[i].Likes=make( map[uint]*User )
	}
}
func (this *User)Getlikes()[]uint{
	result:=make([]uint,len(this.Likes))
	i:=0
	for k,_:=range this.Likes{
		result[i]=k
		i++
	}
	return result
}
func (this *User)Getfans()[]uint{
	result:=make([]uint,len(this.Fans))
	i:=0
	for k,_:=range this.Fans{
		result[i]=k
		i++
	}
	return result
}
func main() {

	http.HandleFunc("/like", func(w http.ResponseWriter,r *http.Request){
		m,_:=url.ParseQuery(r.URL.RawQuery)
		vid,_:=strconv.Atoi(m["vid"][0])
		switch r.Method {
		case http.MethodGet:{
			w.Write([]byte(fmt.Sprint(UserArray[vid].Getlikes())))
			}
		case http.MethodPost:
			star,_:=strconv.Atoi(m["star"][0])
			UserArray[vid].Likes[uint(star)]=UserArray[uint(star)]
			UserArray[uint(star)].Fans[uint(vid)]=UserArray[uint(vid)]
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
			UserArray[vid].Fans[uint(fan)]=UserArray[uint(fan)]
			UserArray[uint(fan)].Likes[uint(vid)]=UserArray[uint(vid)]
		}

	})

	fmt.Println("start http server")
	err :=http.ListenAndServe(":8010",nil)
	panic(err)
}
