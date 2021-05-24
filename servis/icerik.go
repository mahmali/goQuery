package servis

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func Icerik() {
	// Request the HTML page. Html sayfasını isteyin.
	res, err := http.Get("http://islamilimleri.com/Ktphn/Kitablar/05/001/Turkce/01/001.htm")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document. Html belgesini yükleme
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items. İnceleme öğelerini bulun
	doc.Find("#icerik .MsoNormal").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("span").Text()
		//title := s.Find("i").Text()
		/*var yeni string
		iconv.Convert(band, yeni, "windows-1256", "utf-8")*/
		fmt.Printf(band)
	})
}
