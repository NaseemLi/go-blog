// models/banner_model.go
package models

type BannerModel struct {
	Model
	Show  bool   `json:"show"`  //是否展示
	Cover string `json:"cover"` // 图片链接
	Href  string `json:"href"`  // 跳转链接
	Type  int8   `json:"type"`  // 类型 1 banner 2 独家推广
}
