package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetDocumentFromURL(URL string) *goquery.Document {
	content, err := GetBodyString(URL)
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(content)
	if err != nil {
		panic(err)
	}
	return doc
}

//GetBodyString แปลง http body เป็น format utf-8 และ คัดลอกไปยังreader
func GetBodyString(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	body := strings.NewReader(buf.String())
	return body, nil
}
func main() {
	doc := GetDocumentFromURL("http://www.niceoppai.net/onepiece/911/?all")
	doc.Find("#sct_content img").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("src")
		_, found1 := s.Attr("style")
		_, found2 := s.Attr("id")
		if !found1 && !found2 {
			fmt.Printf("Review %d: %s\n", i, url)
		}

	})

}
