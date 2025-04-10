package file

import (
	"errors"
	"goblog/global"
	"goblog/utils"
	"strings"
)

func ImageSuffixJudge(filename string) (suffix string, err error) {
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		err = errors.New("错误的文件名")
		return
	}
	suffix = strings.ToLower(parts[len(parts)-1])

	// 假设 whiteList 是 []string 类型
	if !utils.InList(suffix, global.Config.Upload.WhiteList) {
		err = errors.New("文件非法")
		return
	}
	return
}
