package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	reader, err := os.Open("uploads/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	selection := doc.Find("title").AppendHtml("<meta name=\"keyword1\" content=\"golang-blog\">")
	//fmt.Println(selection.Text())
	selection.SetText("去而且")
	selection.SetAttr("", "")
	html, err := doc.Html()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(html)
}
