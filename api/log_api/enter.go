package logapi

import (
	"fmt"
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/models"
	"goblog/models/enum"
	logservice "goblog/service/log_service"

	"github.com/gin-gonic/gin"
)

type LogApi struct {
}

type LogListRequest struct {
	common.PageInfo
	LogType     enum.LogType      `form:"logType"` // 日志类型 1 2 3
	Level       enum.LogLevelType `form:"level"`   // 日志级别 1 2 3
	UserID      uint              `form:"userID"`  // 用户id
	IP          string            `form:"ip"`
	LoginStatus bool              `form:"loginStatus"`
	ServiceName string            `form:"serviceName"`
}

type LogListResponse struct {
	models.LogModel
	UserNickname string `json:"userNickname"`
	UserAvatar   string `json:"userAvatar"`
}

func (LogApi) LogListView(c *gin.Context) {
	//分页查询 精确查询,模糊匹配
	var cr LogListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	list, count, err := common.ListQuery(models.LogModel{}, common.Options{
		PageInfo:     cr.PageInfo,
		Likes:        []string{"title"},
		Preloads:     []string{"UserModel"},
		DefaultOrder: "created_at desc",
	})

	var _list = make([]LogListResponse, 0)
	for _, logModel := range list {
		_list = append(_list, LogListResponse{
			LogModel:     logModel,
			UserNickname: logModel.UserModel.Nickname,
			UserAvatar:   logModel.UserModel.Avatar,
		})
	}

	res.OkWithList(_list, int(count), c)
	return
}

func (LogApi) LogReadView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var log models.LogModel
	err = global.DB.Take(&log, cr.ID).Error
	if err != nil {
		res.FailWithMsg("不存在的日志", c)
		return
	}

	if !log.IsRead {
		global.DB.Model(&log).Update("is_read", true)
	}

	res.OkWithMsg("日志读取成功", c)
}

func (LogApi) LogRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	log := logservice.Getlog(c)
	log.ShowRequest()
	log.ShowResponse()

	var LogList []models.LogModel
	global.DB.Find(&LogList, "id in ?", cr.IDList)

	if len(LogList) > 0 {
		global.DB.Delete(&LogList)
	}

	msg := fmt.Sprintf("日志删除成功,共删除%d条", len(LogList))

	res.OkWithMsg(msg, c)
}
