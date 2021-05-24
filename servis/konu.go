package servis

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func Konu() {
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
	doc.Find("#icerik > div > table > tbody > tr > td > p:nth-child(1)").Each(func(i int, selection *goquery.Selection) {
		var konu string
		konu = selection.Text()

		fmt.Printf(konu)
	})
}
