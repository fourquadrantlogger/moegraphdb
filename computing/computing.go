package computing

import (
	"encoding/json"
	"fmt"
	"github.com/timeloveboy/moegraphdb/graphdb"
	"runtime"
	"sync"
)

var Start = false
var Now_vid = 1
var Size = 0
var task chan uint = make(chan uint, 10000)
var result chan map[uint]int = make(chan map[uint]int, 1000)

var lock sync.RWMutex
var Result map[uint]int = make(map[uint]int)

func JsonResult() string {
	lock.Lock()
	bs, _ := json.Marshal(Result)
	defer lock.Unlock()
	return string(bs)
}

func Mapper(this graphdb.RelateGraph) {
	fmt.Println("start mapping")
	Size = this.Users.Size()
	for i := 0; i < runtime.NumCPU(); i++ {
		go re(this)
	}
	fmt.Println("start jobber")
	go func() {
		for v := range this.Users.IterItems() {
			task <- v.Key
		}
	}()
	fmt.Println("start duce")
	go func() {
		for true {
			ducer()
		}
	}()
}
func re(this graphdb.RelateGraph) {
	vid := <-task
	u := this.GetUser(vid)
	vid_likes := u.Getlikes()
	vid_likes_max1000000 := this.Filterusers_fanscount(vid_likes, 100*10000, 0)
	count_count := this.GetThemCommonFans(vid_likes_max1000000...)
	count_count_10 := graphdb.Filtercount_min(count_count, 10, 1<<32)
	result <- count_count_10
}
func ducer() {
	c := <-result
	fmt.Println("result", c)
	Now_vid++
	lock.Lock()
	for k, v := range c {
		Result[k] += v
	}
	lock.Unlock()
}
