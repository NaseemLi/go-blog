package conf

import "goblog/conf/site"

type Site struct {
	SiteInfo   site.SiteInfo   `yaml:"siteInfo" json:"siteInfo"`
	Project    site.Project    `yaml:"project" json:"project"`
	Seo        site.Seo        `yaml:"seo" json:"seo"`
	Login      site.Login      `yaml:"login" json:"login"`
	IndexRight site.IndexRight `yaml:"indexRight" json:"indexRight"`
	Article    site.Article    `yaml:"article" json:"article"`
	About      site.About      `yaml:"about" json:"about"`
}
