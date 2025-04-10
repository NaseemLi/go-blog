package main

import (
	"context"
	"fmt"
	"goblog/core"
	"goblog/flags"
	"goblog/global"
	file2 "goblog/utils/file"
	"goblog/utils/hash"
	"io"
	"os"

	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
)

func SendFile(file string) (url string, err error) {
	mac := credentials.NewCredentials(global.Config.QiNiu.AccessKey, global.Config.QiNiu.SecretKey)
	hashString, err := hash.FileMD5(file)
	if err != nil {
		return
	}

	suffix, _ := file2.ImageSuffixJudge(file)
	fileName := fmt.Sprintf("%s.%s", hashString, suffix)
	key := fmt.Sprintf("%s/%s", global.Config.QiNiu.Prefix, fileName)
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	err = uploadManager.UploadFile(context.Background(), file, &uploader.ObjectOptions{
		BucketName: global.Config.QiNiu.Bucket,
		ObjectName: &key,
		FileName:   fileName,
	}, nil)
	return fmt.Sprintf("%s/%s", global.Config.QiNiu.Uri, key), err
}

func SendReader(reader io.Reader) (url string, err error) {
	mac := credentials.NewCredentials(global.Config.QiNiu.AccessKey, global.Config.QiNiu.SecretKey)

	uid := uuid.New().String()
	fileName := fmt.Sprintf("%s.png", uid)
	key := fmt.Sprintf("%s/%s", global.Config.QiNiu.Prefix, fileName)
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	err = uploadManager.UploadReader(context.Background(), reader, &uploader.ObjectOptions{
		BucketName: global.Config.QiNiu.Bucket,
		ObjectName: &key,
		FileName:   fileName,
	}, nil)
	return fmt.Sprintf("%s/%s", global.Config.QiNiu.Uri, key), err
}

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	// url, err := SendFile("uploads/images/WechatIMG33.jpg")
	// fmt.Println(url, err)
	file, _ := os.Open("uploads/images/WechatIMG33.jpg")
	url, err := SendReader(file)
	fmt.Println(url, err)
}
