package commentapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	rediscomment "goblog/service/redis_service/redis_comment"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentDiggView(c *gin.Context) {
	cr := middleware.GetBind[models.IDRequest](c)

	var comment models.CommentModel
	err := global.DB.Take(&comment, cr.ID).Error
	if err != nil {
		res.FailWithMsg("评论不存在", c)
		return
	}
	//查看之前有没有点过赞
	claims := jwts.GetClaims(c)
	//查看之前有没有点过赞
	var userDiggComment models.CommentDiggModel
	err = global.DB.Take(&userDiggComment, "user_id = ? and comment_id = ?", claims.UserID, comment.ID).Error
	if err != nil {
		//点赞
		err = global.DB.Create(&models.CommentDiggModel{
			UserID:    claims.UserID,
			CommentID: cr.ID,
		}).Error
		if err != nil {
			res.FailWithMsg("点赞失败", c)
			return
		}
		rediscomment.SetCacheDigg(comment.ID, 1)
		res.OkWithMsg("点赞成功", c)
		return
	}
	//取消点赞
	global.DB.Delete(&userDiggComment, "user_id = ? and comment_id = ?", claims.UserID, comment.ID)
	res.OkWithMsg("取消点赞成功", c)
	rediscomment.SetCacheDigg(comment.ID, -1)
}
