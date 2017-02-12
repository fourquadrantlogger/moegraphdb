package main

import (
	"fmt"
	"github.com/timeloveboy/moegraphdb/graphdb"
	"github.com/timeloveboy/moegraphdb/routers"
	"net/http"
)

func StaticServer(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/output", http.FileServer(http.Dir("output"))).ServeHTTP(w, r)
}
func main() {
	routers.UserArray = graphdb.NewDB()

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/output/", StaticServer)
	serveMux.HandleFunc("/like", routers.Handle_like)
	serveMux.HandleFunc("/like/n", routers.Handle_like_n)

	serveMux.HandleFunc("/fans", routers.Handle_fans)
	serveMux.HandleFunc("/fans/n", routers.Handle_fans_n)

	serveMux.HandleFunc("/user", routers.Handle_user)

	serveMux.HandleFunc("/relate", routers.Handle_relate)
	serveMux.HandleFunc("/relate/n", routers.Handle_relate_n)

	serveMux.HandleFunc("/common/2/likes", routers.Handle_common_2_like)
	serveMux.HandleFunc("/common/2/fans", routers.Handle_common_2_fans)
	serveMux.HandleFunc("/common/n/like", routers.Handle_common_n_like)
	serveMux.HandleFunc("/common/n/fans", routers.Handle_common_n_fans)

	serveMux.HandleFunc("/count/relate", routers.Handle_count_relate)

	serveMux.HandleFunc("/count/user", routers.Handle_count_user)

	serveMux.HandleFunc("/count/like", routers.Handle_count_like)
	serveMux.HandleFunc("/count/fans", routers.Handle_count_fans)
	serveMux.HandleFunc("/computing/deadfans", routers.Handle_computing_deadfans)
	serveMux.HandleFunc("/computing/autocomputing", routers.AutoComputing)

	fmt.Println("start http server")
	err := http.ListenAndServe(":8010", serveMux)
	if err != nil {
		panic(err)
	}
}
