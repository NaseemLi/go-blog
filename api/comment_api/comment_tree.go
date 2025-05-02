package commentapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	relationshipenum "goblog/models/enum/relationship_enum"
	commentservice "goblog/service/comment_service"
	focusservice "goblog/service/focus_service"
	"goblog/utils"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentTreeView(c *gin.Context) {
	var cr = middleware.GetBind[models.IDRequest](c)

	var article models.ArticleModel
	err := global.DB.Take(&article, "id = ? and status = ?", cr.ID, enum.ArticleStatusPublished).Error
	if err != nil {
		res.FailWithMsg("文章不存在或已被删除", c)
		return
	}

	var userRelationMap = map[uint]relationshipenum.Relation{}
	var userDiggCommentMap = map[uint]bool{}
	claims, err := jwts.ParseTokenByGin(c)
	if err == nil && claims != nil {
		var commentList []models.CommentModel
		global.DB.Find(&commentList, "article_id = ?", cr.ID)
		if len(commentList) > 0 {
			//查我点赞的评论 id 列表
			var commentIDList []uint
			var userIDList []uint
			for _, v := range commentList {
				commentIDList = append(commentIDList, v.ID)
				userIDList = append(userIDList, v.UserID)
			}
			userIDList = utils.Unique(userIDList) //去重
			userRelationMap = focusservice.CalcUserPatchRelationship(claims.UserID, userIDList)
			var commentDiggList []models.CommentDiggModel
			global.DB.Find(&commentDiggList, "user_id = ? and comment_id in ?", claims.UserID, commentIDList)
			for _, v := range commentDiggList {
				userDiggCommentMap[v.CommentID] = true
			}
		}
	}

	//查出根评论
	var commentList []models.CommentModel
	global.DB.Order("created_at desc").Find(&commentList, "article_id = ? and parent_id is null", cr.ID)
	var list = make([]commentservice.CommentResponse, 0)
	for _, model := range commentList {
		response := commentservice.GetCommentTreeV3(model.ID, userRelationMap, userDiggCommentMap)
		list = append(list, *response)
	}

	res.OkWithList(list, len(list), c)
}
