package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	kr "github.com/jofo8948/kristianstad-tester"
)

func main() {
	http.HandleFunc("/data", func (w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		rs := kr.ResultSet{}
		dec.Decode(&rs)
		fmt.Println(rs)
	})

	http.ListenAndServe(":80", nil)
}
