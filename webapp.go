package main

import (
	"fmt"
	"github.com/timeloveboy/moegraphdb/graphdb"
	"github.com/timeloveboy/moegraphdb/routers"
	"net/http"
)

func main() {
	routers.UserArray = graphdb.NewDB()
	http.HandleFunc("/like", routers.Handle_like)
	http.HandleFunc("/like/n", routers.Handle_like_n)

	http.HandleFunc("/fans", routers.Handle_fans)
	http.HandleFunc("/fans/n", routers.Handle_fans_n)

	http.HandleFunc("/user", routers.Handle_user)

	http.HandleFunc("/relate", routers.Handle_relate)
	http.HandleFunc("/relate/n", routers.Handle_relate_n)

	http.HandleFunc("/common/2/likes", routers.Handle_common_2_like)
	http.HandleFunc("/common/2/fans", routers.Handle_common_2_fans)
	http.HandleFunc("/common/n/like", routers.Handle_common_n_like)
	http.HandleFunc("/common/n/fans", routers.Handle_common_n_fans)

	http.HandleFunc("/count/relate", routers.Handle_count_relate)

	http.HandleFunc("/count/user", routers.Handle_count_user)

	http.HandleFunc("/count/like", routers.Handle_count_like)
	http.HandleFunc("/count/fans", routers.Handle_count_fans)
	http.HandleFunc("/computing/deadfans", routers.Handle_computing_deadfans)
	http.HandleFunc("/computing/autocomputing", routers.AutoComputing)
	fmt.Println("start http server")
	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		panic(err)
	}
}
