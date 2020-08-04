package model

import (
	"github.com/zhouya0/blog/pkg/app"
)

type Article struct {
	*Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"convert"`
	CoverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
}

type ArticleSwagger struct {
	List []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}
