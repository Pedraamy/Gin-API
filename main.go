package main

import (

	"github.com/gin-gonic/gin"
	"github.com/pedraamy/gin-api/repository"
	"github.com/pedraamy/gin-api/api"
	"github.com/pedraamy/gin-api/controller"
	"github.com/pedraamy/gin-api/docs"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)


func main() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Go Infra API"
	docs.SwaggerInfo.Description = "Preliminary API for infrastructure provisioning."
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	resourseRepo := repository.NewResourceRepo()
	resourceController := controller.NewController(resourseRepo)
	resourceApi := *api.NewResourceApi(resourceController)

	server := gin.Default()


	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{	
		
		aws := apiRoutes.Group("/aws")
		{
			aws.GET("", resourceApi.GetAwsResources)
			aws.POST("", resourceApi.AddAwsResource)
			//aws.DELETE(":id", videoAPI.DeleteVideo).Use(middlewares.Auth())
		}

		azure := apiRoutes.Group("/azure")
		{
			azure.GET("", resourceApi.GetAzureResources)
			azure.POST("", resourceApi.AddAzureResource)
			//azure.DELETE(":id", videoAPI.DeleteVideo).Use(middlewares.Auth())
		}

		gcp := apiRoutes.Group("/gcp")
		{
			gcp.GET("", resourceApi.GetGcpResources)
			gcp.POST("", resourceApi.AddGcpResource)
			//gcp.DELETE(":id", videoAPI.DeleteVideo).Use(middlewares.Auth())
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := "8080"
	server.Run(":" + port)
}
