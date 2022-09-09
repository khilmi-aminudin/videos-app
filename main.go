package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/videos-app/controller"
	"github.com/khilmi-aminudin/videos-app/service"
	"github.com/khilmi-aminudin/videos-app/utils"
)

var (
	videoService    service.VideoService       = service.NewVideoService()
	videoController controller.VideoController = controller.NewVideoController(videoService)
)

func main() {
	utils.SetLogOutputFile()
	// server := gin.New()
	// server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())
	server := gin.Default()

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "Hello, I am OK!!",
		})
	})

	video := server.Group("/api/video")

	video.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	video.POST("/", func(ctx *gin.Context) {
		video, err := videoController.Save(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "has an error on inputing video",
				"error":   err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, video)
	})

	server.Run(":8080")
}
