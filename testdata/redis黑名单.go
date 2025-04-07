package main

import (
	"fmt"
	"goblog/core"
	"goblog/flags"
	"goblog/global"
	redisjwt "goblog/service/redis_service/redis_jwt"
	"goblog/utils/jwts"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.Redis = core.InitRedis()

	token, err := jwts.GetToken(jwts.Claims{
		UserID: 2,
		Role:   1,
	})
	fmt.Println(token, err)
	redisjwt.TokenBlack(token, redisjwt.UserBlackType)
	blackType, ok := redisjwt.HasTokenBlack(token)
	fmt.Println(blackType, ok)

}
