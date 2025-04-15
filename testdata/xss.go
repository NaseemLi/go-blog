package main

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var md = `# header

# 这是一级标题
## 这是二级标题

Sample text.

[link](http://example.com)
<img src="x" onerror="alert(2)" alt="">
<script>alert(123)</script>
`

	contentDoc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(md)))
	if err != nil {
		return
	}
	contentDoc.Find("script").Remove()
	contentDoc.Find("img").Remove()
	contentDoc.Find("iframe").Remove()

	Content := contentDoc.Text()
	fmt.Println(Content)
}
