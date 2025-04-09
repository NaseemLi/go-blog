package main

import (
	"fmt"

	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/region"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
)

func main() {
	options := uploader.UploadManagerOptions{
		Options: http_client.Options{
			Regions: region.GetRegionByID("z0", true),
		},
	}
	fmt.Println(options)
}
