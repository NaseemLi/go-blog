package main

import (
	"bytes"
	"fmt"
	"goblog/utils/markdown"

	"github.com/PuerkitoBio/goquery"
)

var md = `# header

# 这是一级标题
## 这是二级标题

Sample text.

[link](http://example.com)
`

func main() {
	html := markdown.MdToHTML(md)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
	if err != nil {
		fmt.Println(err)
		return
	}

	htmlText := doc.Text()
	fmt.Println(htmlText)
}
