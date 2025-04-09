package imageapi

import (
	"errors"
	"fmt"
	"goblog/common/res"
	"goblog/global"
	"goblog/models"
	"goblog/utils"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	s := global.Config.Upload.Size
	if fileHeader.Size > s*1024*1024 {
		res.FailWithMsg(fmt.Sprintf("文件大小大于%dMB", s), c)
	}
	//后缀判断
	filename := fileHeader.Filename
	err = imageSuffixJudge(filename)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	//重复判断 hash
	file, err := fileHeader.Open()
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	byteData, _ := io.ReadAll(file)
	hash := utils.Md5(byteData)
	//判断 hash 有无
	var model models.ImageModel
	err = global.DB.Take(&model, "hash = ?", hash).Error
	if err == nil {
		// 找到了
		logrus.Infof("上传图片重复 %s<==>%s %s", filename, model.Filename, hash)
		res.Ok(model.WebPath(), "上传成功", c)
		return
	}
	fmt.Println(hash)
	filePath := fmt.Sprintf("uploads/images/%s/%s", global.Config.Upload.UploadDir, fileHeader.Filename)

	//入库
	model = models.ImageModel{
		Filename: filename,
		Path:     filePath,
		Size:     fileHeader.Size,
		Hash:     hash,
	}
	err = global.DB.Create(&model).Error
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	c.SaveUploadedFile(fileHeader, filePath)
	res.Ok(model.WebPath(), "图片上传成功", c)
}

func imageSuffixJudge(filename string) (err error) {
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		return errors.New("错误的文件名")
	}
	suffix := strings.ToLower(parts[len(parts)-1])

	// 假设 whiteList 是 []string 类型
	if !utils.InList(suffix, global.Config.Upload.WhiteList) {
		return errors.New("文件非法")
	}
	return nil
}
