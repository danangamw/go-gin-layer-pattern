package router

import (
	"go-gin/config"
	"go-gin/handler"
	"go-gin/middleware"
	"go-gin/repository"
	"go-gin/service"

	"github.com/gin-gonic/gin"
)

func PostRouter(api *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	r := api.Group("/tweets")
	r.Use(middleware.JWTMiddleware())
	r.POST("/", postHandler.Create)
}
