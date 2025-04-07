package main

import (
	"fmt"
	"goblog/core"
	"goblog/flags"
	"goblog/global"
	redisjwt "goblog/service/redis_service/redis_jwt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.Redis = core.InitRedis()

	// token, err := jwts.GetToken(jwts.Claims{
	// 	UserID: 2,
	// 	Role:   1,
	// })
	// fmt.Println(token, err)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjEsInVzZXJuYW1lIjoiIiwicm9sZSI6MSwiZXhwIjoxNzQ0MDM5NDU1LCJpc3MiOiJuYXNlZW0ifQ.WsVTCtTvJS4zAhC3IDN-UmpU6EpZUQE2vBSiREGcVpw"
	redisjwt.TokenBlack(token, redisjwt.UserBlackType)
	blackType, ok := redisjwt.HasTokenBlack(token)
	fmt.Println(blackType, ok)

}
