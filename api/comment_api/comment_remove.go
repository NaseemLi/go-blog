package commentapi

import (
	"fmt"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	commentservice "goblog/service/comment_service"
	rediscomment "goblog/service/redis_service/redis_comment"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentRemoveView(c *gin.Context) {
	var cr = middleware.GetBind[models.IDRequest](c)

	claims := jwts.GetClaims(c)

	//普通用户只能删自己发的评论和自己文章下的评论
	var model models.CommentModel
	err := global.DB.Preload("ArticleModel").Take(&model, cr.ID).Error
	if err != nil {
		res.FailWithMsg("评论不存在或已被删除", c)
		return
	}

	if claims.Role != enum.AdminRole {
		if !(model.UserID == claims.UserID || model.ArticleModel.UserID == claims.UserID) {
			res.FailWithMsg("没有权限删除该评论", c)
			return
		}
	}

	subList := commentservice.GetCommentOneDimensionalTree(model.ID)
	//删除评论,找所有的子评论,还要找所有的父评论
	if model.ParentID != nil {
		//查找父评论
		parentList := commentservice.GetParents(*model.ParentID)
		for _, v := range parentList {
			rediscomment.SetCacheApply(v.ID, -len(subList))
		}
	}
	//删掉评论
	global.DB.Delete(&subList)

	res.OkWithMsg(fmt.Sprintf("评论删除成功,共删除%d条评论", len(subList)), c)
}
