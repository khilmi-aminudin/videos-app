package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/videos-app/entity"
	"github.com/khilmi-aminudin/videos-app/service"
)

type VideoController interface {
	Save(ctx *gin.Context) (entity.Video, error)
	FindAll() []entity.Video
}

type videoController struct {
	service service.VideoService
}

func NewVideoController(service service.VideoService) VideoController {
	return &videoController{
		service: service,
	}
}

// Save implements videoController
func (c *videoController) Save(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	err := ctx.BindJSON(&video)

	if err != nil {
		return video, err
	}

	c.service.Save(video)

	return video, nil
}

// FindAll implements videoController
func (c *videoController) FindAll() []entity.Video {
	return c.service.FindAll()
}
