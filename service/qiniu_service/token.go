package qiniuservice

import (
	"context"
	"goblog/global"
	"time"

	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
)

func GenToken() (token string, err error) {
	mac := credentials.NewCredentials(global.Config.QiNiu.AccessKey, global.Config.QiNiu.SecretKey)
	bucket := global.Config.QiNiu.Bucket
	putPolicy, err := uptoken.NewPutPolicy(bucket, time.Now().Add(time.Duration(global.Config.QiNiu.Expiry)*time.Second))
	if err != nil {
		return
	}
	token, err = uptoken.NewSigner(putPolicy, mac).GetUpToken(context.Background())
	if err != nil {
		return
	}
	return
}
