package main

import (
	"fmt"
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
	MaxTests = 30
)

var (
urls = []string {
		"http://www.kristianstad.se",
		"https://www.kristianstad.se",
		"https://www.kristianstad.se/sv/e-tjanster/",
		"https://www.kristianstad.se/globalassets/blanketter/barn-och-utbildning/ansokan_tillaggsbelopp_frst.pdf",
		"https://www.kristianstad.se/sv/barn-och-utbildning/grundskola/",
		"https://www.kristianstad.se/sv/bygga-bo-och-miljo/bygga-nytt-andra-eller-riva/bygglov/",
		"https://www.kristianstad.se/sv/bygga-bo-och-miljo/bostader/hitta-bostad/",
		"https://www.kristianstad.se/sv/huvudnyheter/",
		"https://www.kristianstad.se/sv/trafik-och-resor/trafik-resor-och-gator/",
		"https://www.kristianstad.se/sv/kommun-och-politik/overklaga-beslut-rattssakerhet/",
		"http://turism.kristianstad.se/",
		"http://" + centralServer,
	}

	bar = pb.New(MaxTests*len(urls))
)

func main() {
	rs := test()
	sendToServer(rs)
	shutdown()
}

func test() kr.ResultSet {
	rs := &kr.ResultSet{}
	log.SetOutput(rs)
	defer log.SetOutput(os.Stdout)
	var err error
	rs.Name, err = windows.ComputerName()
	if err != nil {
		log.Fatal("Could not get name of computer.")
	}

	rs.User = os.Getenv("USERNAME")

	ticker := time.NewTicker(1*time.Minute);
	rs.StartTime = time.Now();

	bar.ShowTimeLeft = false
	bar.ShowBar = true
	bar.Start()

loop:
	for i := 0; i < MaxTests; i++ {
		select {
			case <-ticker.C:
					res := runTest(urls, i)
					rs.Results = append(rs.Results, res...)
			case <-time.After(3*time.Hour):
				ticker.Stop(); break loop
		}
	}

	rs.EndTime = time.Now()
	bar.FinishPrint("Klart!")
	return *rs;
}

func runTest(urls []string, iter int) (rs []kr.Result) {
	for _, url := range urls {
		r := testUrl(url)
		r.Iteration = iter
		rs = append(rs, r)
		bar.Increment()
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

func sendToServer(rs kr.ResultSet) {
	fmt.Println("Skickar den insamlade informationen till Sigma ITC...")
	rd, wrt := io.Pipe();

	go func() {
		defer wrt.Close()
		enc := json.NewEncoder(wrt)
		enc.Encode(rs)
	}()

	_, err := http.Post("http://" + centralServer + "/data", "application/json", rd)
	if err != nil {
		log.Fatalln(err)
	}
}

func shutdown() {
	fmt.Println("Klart. Programmet kommer nu stänga sig självt.")
	fmt.Println("Ha en fortsatt bra dag!")
	time.Sleep(3*time.Second)
	os.Exit(0)
}
