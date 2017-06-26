package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"database/sql"
	_ "github.com/lib/pq"
	kr "github.com/jofo8948/kristianstad-tester"
)

func main() {
	http.HandleFunc("/data", func (w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		rs := kr.ResultSet{}
		dec.Decode(&rs)
		fmt.Println(rs)
		db, err := sql.Open("postgres","user=perf password=RYecS6vt dbName=performace sslmode=disable")
		if err != nil {
			log.Fatal("Error: databasargumenten är felaktiga", err)
		}

		defer db.Close();

		err = db.Ping()
		if err != nil {
			log.Fatal("Error: Kunde inte ansluta till DB", err)
		}


		qs, err := db.Prepare("INSERT INTO ResultSet (name, start_date, end_date) VALUES ($1,$2,$3) RETURNING id;")
		if err != nil {
		 	log.Fatal(err)
		}

		var inserted_id int;
		err = qs.QueryRow(rs.Name, rs.StartTime, rs.EndTime).Scan(&inserted_id)
		if err != nil {
			log.Fatal("Could not store ResultSet in DB.", err)
		}

		rqs, err := db.Prepare("INSERT INTO Result (url, start_date, duration, resultset) VALUES ($1,$2,$3,$4);")
		for _, x := range rs.Results {
			rqs.Exec(x.Url, x.StartTime, x.Duration, inserted_id)
		}

	})

	http.ListenAndServe(":80", nil)
}
