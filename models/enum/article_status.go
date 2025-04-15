package enum

type ArticleStatusType int8

const (
	ArticleStatusDraft     ArticleStatusType = 1 //草稿
	ArticleStatusExamine   ArticleStatusType = 2 //审核中
	ArticleStatusPublished ArticleStatusType = 3 //已发布
)
