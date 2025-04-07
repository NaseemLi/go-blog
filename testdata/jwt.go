package main

import (
	"fmt"
	"goblog/core"
	"goblog/flags"
	"goblog/global"
	"goblog/utils/jwts"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	token, err := jwts.GetToken(jwts.Claims{
		UserID: 1,
		Role:   2,
	})
	fmt.Println(token, err)

	// cls, err2 := jwts.ParseToken(token)
	// fmt.Println(cls, err2)
}

// func ParseTokenByGin(c *gin.Context) (*MyClaims, error) {
// 	token := c.GetHeader("token")
// 	if token == "" {
// 		token = c.Query("token")
// 	}
// 	return ParseToken(token)
// }
