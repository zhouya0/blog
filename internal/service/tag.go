package service

type CountTagRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

