package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	var postdatalimit int = 1000000

	db, err := sql.Open("mysql", "paidian:Paidian2016@tcp(rm-2ze076baorsp7ed0lo.mysql.rds.aliyuncs.com:3306)/crawler_meipai")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
		panic(nil)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	for idstart := 0; ; idstart += postdatalimit {
		relate := make([]struct {
			Vid1   uint `json:"vid1"`
			Vid2   uint `json:"vid2"`
			Relate int  `json:"relate"`
		}, postdatalimit)

		rows, err := db.Query("select user_id,fans_id from meipai_rapport_t1 limit ? offset ?", postdatalimit, idstart)
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()
		i := 0
		for rows.Next() {
			var user_id, fans_id uint
			err := rows.Scan(&user_id, &fans_id)
			if err != nil {
				panic(err)
			}
			relate[i] = struct {
				Vid1   uint `json:"vid1"`
				Vid2   uint `json:"vid2"`
				Relate int  `json:"relate"`
			}{
				Vid1:   user_id,
				Vid2:   fans_id,
				Relate: 2,
			}
			i++
		}
		err = rows.Err()
		if err != nil {
			panic(err)
		}

		bs, err := json.Marshal(relate)
		//post(string(bs))
		fmt.Print(time.Now().String()[:19], "\t")
		fmt.Println(idstart, len(bs))
	}
}
