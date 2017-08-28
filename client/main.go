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
	MaxTests = 5
	TimeBetweenTests = 1 * time.Minute
)

var (
urls = []string {
		"http://" + centralServer,
		"https://intranat.kristianstad.se",
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

		"https://fonts.googleapis.com/css?family=Open+Sans:400,300,600,700,800,600italic,400italic,300italic",
		"https://www.kristianstad.se/Static/css/EPiCoreCss?v=PWgx8JbL-aPuKIA3UEqkkQcLAd6GU4nLbYYomLaPXBM1",
		"https://www.kristianstad.se/Static/css/ExtraCss?v=Trv3sRusoE6hwzOnF6xPR4wpA6RslkV7h0mBWkrLo941",
		"https://www.kristianstad.se/Static/img/logo/logo.png",
		"https://www.kristianstad.se/contentassets/c49d3ed14d9b4a5191be6ba5dfab2553/lingenas.jpg?preset=page-puff",
		"https://www.kristianstad.se/globalassets/bilder/barn-och-utbildning/musikskolan/nisse_1.jpg?preset=page-puff",
		"https://www.kristianstad.se/globalassets/bilder/uppleva-och-gora/barbacka/fyrkant-huset-2.jpg?preset=page-puff",
		"https://www.kristianstad.se/globalassets/bilder/uppleva-och-gora/barbacka/fyrkant-off-program_sida_1.jpg?preset=page-puff",
		"https://www.kristianstad.se/contentassets/409374eabf684c949df808929e8e4af1/vattentankar_everod.jpg?preset=page-puff",
		"https://www.kristianstad.se/contentassets/b40c7f8d19e940ef9c2f0fadffaaff43/ollsjo_utomhuslangor_webb.jpg?preset=page-puff",
		"https://translate.google.com/translate_a/element.js?cb=googleTranslateElementInit",
		"https://www.kristianstad.se/contentassets/b0c75094593c46979006fd2d454f58f8/krinova_740.jpg?preset=page-puff",
		"https://fonts.gstatic.com/s/opensans/v14/MTP_ySUJH_bn48VBG8sNSugdm0LZdjqr5-oayXSOefg.woff2",
		"https://www.kristianstad.se/Static/font/opensans/Bold/OpenSans-Bold.woff2",
		"https://fonts.gstatic.com/s/opensans/v14/DXI1ORHCpsQm3Vp6mXoaTegdm0LZdjqr5-oayXSOefg.woff2",
		"https://www.kristianstad.se/Static/font/opensans/Regular/OpenSans-Regular.woff2",
		"https://www.kristianstad.se/Static/font/font-awesome/fontawesome-webfont.woff2?v=4.5.0",
		"https://www.kristianstad.se/Static/img/grey_video_poster.png",
		"https://www.kristianstad.se/contentassets/95d7b796c22945b691088c544417bc38/balsby-badplats.jpg?preset=page-puff",
		"https://www.kristianstad.se/Static/img/icons/icon-kommun.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-uppleva-white.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-utbildningochbarnomsorg-white.png",
		"https://www.kristianstad.se/Static/img/icons/icon-trafikochresor-white.png",
		"https://www.kristianstad.se/Static/img/icons/arrow-down.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-sok.svg",
		"https://www.kristianstad.se/Static/img/menu.png",
		"https://www.kristianstad.se/Static/img/icons/icon-bygga-white.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-naringsliv.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-bygga.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-omsorgochhjalp-white.png",
		"https://www.kristianstad.se/Static/img/icons/icon-trafikochresor.png",
		"https://www.kristianstad.se/Static/img/icons/icon-utbildningochbarnomsorg.png",
		"https://www.kristianstad.se/Static/img/icons/icon-naringsliv-white.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-kommun-white.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-uppleva.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-omsorgochhjalp.png",
		"https://www.kristianstad.se/Static/img/icons/icon-etjanster.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-translate-hc.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-kontakta.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-anpassa.svg",
		"https://www.vizzit.se/j/etcb/?c=kristianstad2017_se&h=JdDAry9XLVe0kZ3oo3J3vA==",
		"https://spoxy4.insipio.com/filepath/www.kristianstad.se/js/speakit.min.js",
		"https://translate.googleapis.com/translate_static/css/translateelement.css",
		"https://translate.googleapis.com/translate_static/js/element/main_sv.js",
		"https://www.kristianstad.se/Static/js/EPiCoreJs?v=3dgVYUxOAX5AmuCdavaaFMaEHSyI1qgdBd2E0kFJZ9g1",
		"https://www.kristianstad.se/Static/js/ExtraJs?v=sQzpRrgFX1EZtxN6gcMjmcuybYD4E2lh_6WslMppnxY1",
		"https://www.kristianstad.se/globalassets/bilder/barn-och-utbildning/musikskolan/audition-flash.jpg?preset=page-puff",
		"https://www.kristianstad.se/globalassets/bilder/logotyper/abk_240x135.jpg?preset=startpage-slideshow",
		"https://www.kristianstad.se/globalassets/bilder/logotyper/c4eneri_240x135.jpg?preset=startpage-slideshow",
		"https://www.kristianstad.se/globalassets/bilder/logotyper/krinova_iosp_rgb_240x135.jpg?preset=startpage-slideshow",
		"https://translate.googleapis.com/element/TE_20170814_01/e/js/element/element_main.js",
		"https://www.kristianstad.se/globalassets/bilder/logotyper/krstd-airport-logo-240x135.jpg?preset=startpage-slideshow",
		"https://www.kristianstad.se/globalassets/bilder/logotyper/renhallningen_240x135.jpg?preset=startpage-slideshow",
		"https://www.kristianstad.se/globalassets/bilder/logotyper/skanenordost_240x135.jpg?preset=startpage-slideshow",
		"https://www.kristianstad.se/globalassets/bilder/logotyper/snoka240x135.jpg?preset=startpage-slideshow",
		"https://www.kristianstad.se/globalassets/bilder/logotyper/vattenriket_250x135.jpg?preset=startpage-slideshow",
		"https://www.kristianstad.se/globalassets/bilder/logotyper/ahus_hamn240x135.jpg?preset=startpage-slideshow",
		"https://www.kristianstad.se/globalassets/bilder/logotyper/kulturkvarteret_kristianstad.jpg?preset=startpage-slideshow",
		"https://www.kristianstad.se/globalassets/webmaster/video-startsida/sommarfilm.mp4",
		"https://www.kristianstad.se/Static/img/icons/arrow-right-green.svg",
		"https://www.kristianstad.se/Static/img/icons/shadow.png",
		"https://www.kristianstad.se/Static/img/footer_bg.png",
		"https://www.kristianstad.se/Static/img/icons/icon-hittasnabbt-green.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-omkommunen-green.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-kontakta-green.svg",
		"https://www.kristianstad.se/Static/img/icons/facebook.svg",
		"https://www.kristianstad.se/Static/img/icons/twitter.svg",
		"https://www.kristianstad.se/Static/img/icons/linkedin.svg",
		"https://www.kristianstad.se/Static/img/icons/flickr.svg",
		"https://www.kristianstad.se/Static/img/icons/youtube.svg",
		"https://www.google-analytics.com/analytics.js",
		"https://www.vizzit.se/vizzittag/assets/pixel.js?0.19000211864582184&d=kristianstad2017_se&r=https%3A%2F%2Fwww.kristianstad.se%2Ffacebook.svg&l=https%3A%2F%2Fwww.kristianstad.se%2Fsv%2F&c=ftYgHpDubpFWc6ONXDOO9g%3D%3D%3A1499069644&sx=1920&sy=1080&bl=sv&fe=false&fv=0,0,0&je=false&u=&anonip=false",
		"https://www.vizzit.se/overlay/episerver_new/overlay.php?c=JdDAry9XLVe0kZ3oo3J3vA==",
		"https://www.vizzit.se/overlay/broken/js/overlay.min.js",
		"https://connect.facebook.net/en_US/sdk.js",
		"https://platform.twitter.com/widgets.js",
		"https://www.kristianstad.se/Static/img/icons/icon-pil.svg",
		"https://www.kristianstad.se/Static/img/icons/icon-laddafler.svg",
		"https://www.google-analytics.com/collect?v=1&_v=j60&a=1534596473&t=pageview&_s=1&dl=https%3A%2F%2Fwww.kristianstad.se%2Fsv%2F&ul=sv&de=UTF-8&dt=Startsida%20-%20Kristianstads%20kommun&sd=24-bit&sr=1920x1080&vp=866x901&je=0&_u=AACAAEABI~&jid=&gjid=&cid=598453190.1499069645&tid=UA-182210-18&_gid=965251074.1503901727&z=1221376550",
		"https://translate.googleapis.com/translate_a/l?client=te&alpha=true&hl=sv&cb=_callbacks____0j6vvjm02",
		"https://translate.googleapis.com/translate_static/css/translateelement.css",
		"https://www.gstatic.com/images/branding/product/1x/translate_24dp.png",
		"https://www.google.com/images/cleardot.gif",
		"https://www.gstatic.com/images/branding/product/2x/translate_24dp.png",
		"https://www.facebook.com/impression.php/f6a0333c71b8d4/?api_key=1803562503209635&lid=115&payload=%7B%22source%22%3A%22jssdk%22%7D",
		"https://www.facebook.com/impression.php/f6a0333c71b8d4/?api_key=1803562503209635&lid=115&payload=%7B%22source%22%3A%22jssdk%22%7D",
		"https://staticxx.facebook.com/connect/xd_arbiter/r/0sTQzbapM8j.js?version=42",
		"https://www.kristianstad.se/globalassets/webmaster/bild/tostebergahamn.jpg",

	}

	bar = pb.New(MaxTests)
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

	deadline := time.After(3*time.Hour)
	ticker := time.NewTicker(TimeBetweenTests);
	rs.StartTime = time.Now();

	bar.ShowTimeLeft = true
	bar.ShowBar = true
	bar.Start()

	{
		res := runTest(urls, 0)
		rs.Results = append(rs.Results, res...)
		bar.Increment()
	}

loop:
	for i := 1; i < MaxTests; i++ {
		select {
			case <-ticker.C:
					res := runTest(urls, i)
					rs.Results = append(rs.Results, res...)
			case <-deadline:
				ticker.Stop(); break loop
		}
		bar.Increment()
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
	r.Duration = time.Since(r.StartTime)

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
