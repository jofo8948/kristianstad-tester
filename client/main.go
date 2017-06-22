package main

import (
	"time"
	"net/http"
	"os"
	"log"
	"io"
	"io/ioutil"
	"encoding/json"
	kr "github.com/jofo8948/kristianstad-tester"
	"golang.org/x/sys/windows"
	"gopkg.in/cheggaaa/pb.v1"
)

const (
	centralServer = "146.185.158.83"
	MaxTests = 2
)

func main() {
	rs := &kr.ResultSet{}
	log.SetOutput(rs)
	defer log.SetOutput(os.Stdout)
	var err error
	rs.Name, err = windows.ComputerName()
	if err != nil {
		log.Fatal("Could not get name of computer.")
	}

	urls := []string {
		"http://www.kristianstad.se",
		"https://www.kristianstad.se",
		"https://www.kristianstad.se/sv/barn-och-utbildning/grundskola/",
		"https://www.kristianstad.se/sv/bygga-bo-och-miljo/bygga-nytt-andra-eller-riva/bygglov/",
		"https://www.kristianstad.se/sv/bygga-bo-och-miljo/bostader/hitta-bostad/",
		"https://www.kristianstad.se/sv/huvudnyheter/",
	}

	ticker := time.NewTicker(1*time.Minute);
	rs.StartTime = time.Now();
	bar := pb.New(MaxTests)
	bar.ShowTimeLeft = false
	bar.ShowBar = true
	bar.Start()

loop:
	for i := 0; i < MaxTests; i++ {
		select {
			case <-ticker.C:
					res := runTest(urls)
					bar.Increment()
					rs.Results = append(rs.Results, res...)
			case <-time.After(3*time.Hour):
				ticker.Stop(); break loop
		}
	}

	rs.EndTime = time.Now()


	rd, wrt := io.Pipe();

	go func() {
		enc := json.NewEncoder(wrt)
		enc.Encode(rs);
		wrt.Close();
	}()
	http.Post(centralServer + "/data", "application/json", rd)
	bar.FinishPrint("Klart!")
	os.Exit(0);
}

func runTest(urls []string) (rs []kr.Result) {
	for _, url := range urls {
		r := testUrl(url)
		rs = append(rs, r)
	}
	return;
}

func testUrl(url string) (r kr.Result) {
	r.Url = url;
	r.StartTime = time.Now()

	res, err := http.Get(url)
	if (err != nil) {
		log.Print(err)
		return
	}

	r.StatusCode = res.StatusCode

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	r.Size = len(data)
	r.Duration = time.Now().Sub(r.StartTime)

	return r;
}
