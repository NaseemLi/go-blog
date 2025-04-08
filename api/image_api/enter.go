package imageapi

import (
	"errors"
	"fmt"
	"goblog/common/res"
	"goblog/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type ImageApi struct {
}

func (ImageApi) ImageUploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	//判断文件大小
	if fileHeader.Size > 2*1024*1024 {
		res.FailWithMsg("文件大小大于 2 MB", c)
	}
	//后缀判断
	filename := fileHeader.Filename
	err = imageSuffixJudge(filename)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	filePath := fmt.Sprintf("uploads/images/%s", fileHeader.Filename)
	c.SaveUploadedFile(fileHeader, filePath)
	res.Ok("/"+filePath, "图片上传成功", c)
}

var whiteList = []string{
	"jpg",
	"jpeg",
	"png",
	"webp",
	"git",
}

func imageSuffixJudge(filename string) (err error) {
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		return errors.New("错误的文件名")
	}
	suffix := parts[len(parts)-1]

	// 假设 whiteList 是 []string 类型
	if !utils.InList[string](suffix, whiteList) {
		return errors.New("文件非法")
	}
	return nil
}
