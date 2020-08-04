package model

import (
	"github.com/zhouya0/blog/pkg/app"
)

type Tag struct {
	*Model
	Name string `json:"name"`
	State uint8 `json:"state"`
}

type TagSwagger struct {
	List []*Tag
	Pager *app.Pager
}

func (a Tag) TableName() string {
	return "blog_tag"
}
