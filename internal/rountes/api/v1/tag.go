package v1

import "github.com/gin-gonic/gin"

type Tag struct {}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) GET (c *gin.Context) {}

func (t Tag) List (c *gin.Context) {}

func (t Tag) Create (c *gin.Context) {}

