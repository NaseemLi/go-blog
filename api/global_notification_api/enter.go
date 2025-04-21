package globalnotificationapi

import (
	"fmt"
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type GlobalNotificationApi struct {
}

type CreateRequest struct {
	Title   string `json:"title" binding:"required"`
	Icon    string `json:"icon"`
	Content string `json:"content" binding:"required"`
	Href    string `json:"href"`
}

func (GlobalNotificationApi) CreateView(c *gin.Context) {
	cr := middleware.GetBind[CreateRequest](c)

	var model models.GlobalNotificationModel
	err := global.DB.Take(&model, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMsg("全局消息名称重复", c)
		return
	}

	err = global.DB.Create(&models.GlobalNotificationModel{
		Title:   cr.Title,
		Icon:    cr.Icon,
		Content: cr.Content,
		Href:    cr.Href,
	}).Error
	if err != nil {
		res.FailWithMsg("全局消息创建失败", c)
		return
	}

	res.OkWithMsg("全局消息创建成功", c)
}

type ListRequest struct {
	common.PageInfo
	Type int8 `json:"type" binding:"required,oneof=1 2"`
}

type ListResponse struct {
	models.GlobalNotificationModel
	IsRead bool `json:"isRead"`
}

func (GlobalNotificationApi) ListView(c *gin.Context) {
	cr := middleware.GetBind[ListRequest](c)

	claims := jwts.GetClaims(c)
	readMsgMap := make(map[uint]bool)
	query := global.DB.Where("")

	switch cr.Type {
	case 1: //没被用户删除的
		var ugnmList []models.UserGlobalNotificationModel
		global.DB.Find(&ugnmList, "user_id = ? and is_delete = ?", claims.UserID, false)
		var msgIDList []uint
		for _, v := range ugnmList {
			readMsgMap[v.NotificationID] = v.IsRead
			msgIDList = append(msgIDList, v.ID)
		}

		query = global.DB.Where("id in (?)", msgIDList)
	case 2:
		if claims.Role != enum.AdminRole {
			res.FailWithMsg("没有权限", c)
			return
		}
	}
	_list, _, _ := common.ListQuery(models.GlobalNotificationModel{}, common.Options{
		PageInfo: cr.PageInfo,
		Likes:    []string{"title", "content"},
		Where:    query,
	})
	list := make([]ListResponse, 0)
	for _, v := range _list {
		list = append(list, ListResponse{
			GlobalNotificationModel: v,
			IsRead:                  readMsgMap[v.ID],
		})
	}

	res.OkWithList(list, len(list), c)
}

func (GlobalNotificationApi) RemoveView(c *gin.Context) {
	cr := middleware.GetBind[models.RemoveRequest](c)

	var list []models.GlobalNotificationModel
	global.DB.Find(&list, "id in ?", cr.IDList)
	if len(list) > 0 {
		global.DB.Delete(&list)
	}
	res.OkWithMsg(fmt.Sprintf("删除全局消息 %d个,成功%d个", len(cr.IDList), len(list)), c)
}

type UserMsgActionViewRequest struct {
	ID   uint `json:"id" binding:"required"`
	Type int8 `json:"type" binding:"required,oneof=1 2"` //1:已读 2:删除
}

func (GlobalNotificationApi) UserMsgActionView(c *gin.Context) {
	cr := middleware.GetBind[UserMsgActionViewRequest](c)

	var msg models.GlobalNotificationModel
	err := global.DB.Take(&msg, "id = ?", cr.ID).Error
	if err != nil {
		res.FailWithMsg("全局消息不存在", c)
		return
	}

	claims := jwts.GetClaims(c)
	model := models.UserGlobalNotificationModel{
		NotificationID: cr.ID,
		UserID:         claims.UserID,
	}

	m := "消息读取成功"
	if cr.Type == 1 {
		model.IsRead = true
	} else {
		model.IsDelete = true
		m = "消息删除成功"
	}
	// 看一看之前有没有操作过
	var ugnm models.UserGlobalNotificationModel
	err = global.DB.Take(&ugnm, "user_id = ? and notification_id = ?", claims.UserID, cr.ID).Error
	// 之前这个用户对这个消息没有操作过
	// 之前对这个消息有读取操作
	// 之前对这个消息有删除操作
	// 先删除再读取
	if err != nil {
		global.DB.Create(&model)
		res.OkWithMsg("消息读取成功", c)
		return
	}

	if ugnm.IsDelete {
		res.FailWithMsg("消息已删除", c)
		return
	}

	if ugnm.IsRead && ugnm.IsDelete {
		res.FailWithMsg("消息已删除", c)
		return
	}

	if ugnm.IsRead {
		//如果现在删除成功就更新
		if model.IsDelete {
			global.DB.Model(&ugnm).Update("is_read", true)
		}
	}

	res.OkWithMsg(m, c)
}
