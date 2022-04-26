package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/api"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/docs" // Swagger generated files
	"gitlab.com/pragmaticreviews/golang-gin-poc/middlewares"
	"gitlab.com/pragmaticreviews/golang-gin-poc/repository"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func main() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Go Infra API"
	docs.SwaggerInfo.Description = "Preliminary API for infrastructure provisioning."
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"https"}

	defer videoRepository.CloseDB()

	server := gin.Default()

	videoAPI := api.NewVideoAPI(loginController, videoController)

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{	
		
		aws := apiRoutes.Group("/aws")
		{
			aws.GET("", videoAPI.GetVideos).Use(middlewares.Auth())
			aws.POST("", videoAPI.CreateVideo).Use(middlewares.Auth())
			aws.DELETE(":id", videoAPI.DeleteVideo).Use(middlewares.Auth())
		}

		azure := apiRoutes.Group("/azure")
		{
			azure.GET("", videoAPI.GetVideos).Use(middlewares.Auth())
			azure.POST("", videoAPI.CreateVideo).Use(middlewares.Auth())
			azure.DELETE(":id", videoAPI.DeleteVideo).Use(middlewares.Auth())
		}

		gcp := apiRoutes.Group("/gcp")
		{
			gcp.GET("", videoAPI.GetVideos).Use(middlewares.Auth())
			gcp.POST("", videoAPI.CreateVideo).Use(middlewares.Auth())
			gcp.DELETE(":id", videoAPI.DeleteVideo).Use(middlewares.Auth())
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := "8080"
	server.Run(":" + port)
}
