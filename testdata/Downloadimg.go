package main

import (
	"fmt"
	"regexp"
)

func main() {
	x := `介绍\n![](https://i1.hdslb.com/bfs/archive/832be5e18a614f7bcd93bf34c20ce690a044d656.jpg@672w_378h_1c_!web-home-common-cover.webp)`

	regex := regexp.MustCompile(`!$begin:math:display$.*?$end:math:display$$begin:math:text$(.*?)$end:math:text$`)

	xx := regex.ReplaceAllStringFunc(x, func(s string) string {
		src := regex.FindStringSubmatch(s)
		fmt.Println(ss)
		return s
	})

	fmt.Println(xx)
}
