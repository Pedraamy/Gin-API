package main

import (

	"github.com/gin-gonic/gin"
	"github.com/pedraamy/gin-api/api"
	"github.com/pedraamy/gin-api/controller"
	"github.com/pedraamy/gin-api/repository"
	"github.com/pedraamy/gin-api/api"
	"github.com/pedraamy/gin-api/controller"
	"github.com/pedraamy/gin-api/middlewares"
	"github.com/pedraamy/gin-api/respository"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var (
	resourseRepo repository.ResourceRepo = repository.NewResourceRepo()
	resourceController controller.ResourceController = controller.NewController(resourseRepo)
	resourceApi api.ResourceApi = *api.NewResourceApi(resourceController)
)

func main() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Go Infra API"
	docs.SwaggerInfo.Description = "Preliminary API for infrastructure provisioning."
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"https"}


	defer resourseRepo.CloseDB()

	server := gin.Default()


	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{	
		
		aws := apiRoutes.Group("/aws")
		{
			aws.GET("", resourceApi.GetAwsResources).Use(middlewares.Auth())
			aws.POST("", resourceApi.AddAwsResource).Use(middlewares.Auth())
			//aws.DELETE(":id", videoAPI.DeleteVideo).Use(middlewares.Auth())
		}

		azure := apiRoutes.Group("/azure")
		{
			azure.GET("", resourceApi.GetAzureResources).Use(middlewares.Auth())
			azure.POST("", resourceApi.AddAzureResource).Use(middlewares.Auth())
			//azure.DELETE(":id", videoAPI.DeleteVideo).Use(middlewares.Auth())
		}

		gcp := apiRoutes.Group("/gcp")
		{
			gcp.GET("", resourceApi.GetGcpResources).Use(middlewares.Auth())
			gcp.POST("", resourceApi.AddGcpResource).Use(middlewares.Auth())
			//gcp.DELETE(":id", videoAPI.DeleteVideo).Use(middlewares.Auth())
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := "8080"
	server.Run(":" + port)
}
