package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/json"
	"database/sql"
	_ "github.com/lib/pq"
	kr "github.com/jofo8948/kristianstad-tester"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Det här är en testsida för Kristianstad-testet. Om du vill veta mer om experimentet, kontakta johan (punkt) fogelstrom (snabel-a) sigma (punkt) se.")
	})

	http.HandleFunc("/data", func (w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		rs := kr.ResultSet{}
		dec.Decode(&rs)
		fmt.Println(rs)

		if (rs.Name == "" || rs.StartTime.Year() < 2017 || rs.EndTime.Year() < 2017) {
			fmt.Fprintln(w, "Ogiltig indata. Kontrollera namn och de angivna tiderna.");
			return;
		}

		user, passw, dbname := os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME")
		connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, passw, dbname)
		db, err := sql.Open("postgres",connectionString)
		if err != nil {
			log.Fatal("Error: databasargumenten är felaktiga", err)
		}

		defer db.Close();

		err = db.Ping()
		if err != nil {
			log.Fatal("Error: Kunde inte ansluta till DB", err)
		}


		qs, err := db.Prepare("INSERT INTO ResultSets (name, start_date, end_date) VALUES ($1,$2,$3) RETURNING id;")
		if err != nil {
		 	log.Fatal(err)
		}

		var inserted_id int;
		err = qs.QueryRow(rs.Name, rs.StartTime, rs.EndTime).Scan(&inserted_id)
		if err != nil {
			log.Fatal("Could not store ResultSet in DB.", err)
		}

		lqs, err := db.Prepare("INSERT INTO Logs (resultset, message) VALUES ($1,$2);")
		for _, x := range rs.Log {
			_, err = lqs.Exec(inserted_id, x)
			if err != nil {
				log.Fatal("Could not store log entry: ", err)
			}
		}

		rqs, err := db.Prepare("INSERT INTO Results (url, comment, start_date, duration, statuscode, size, resultset, iteration) VALUES ($1,$2,$3,$4,$5,$6,$7,$8);")
		for _, x := range rs.Results {
			_, err = rqs.Exec(x.Url, x.Comment, x.StartTime, x.Duration, x.StatusCode, x.Size, inserted_id, x.Iteration)
			if err != nil {
				log.Fatal("Could not store result entry: ", err, x)
			}
		}



		fmt.Fprintln(w, "Tack för ditt bidrag.")
	})

	http.ListenAndServe(":80", nil)
}
