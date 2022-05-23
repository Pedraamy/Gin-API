package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedraamy/gin-api/controller"
	"github.com/pedraamy/gin-api/dto"
)

type ResourceApi struct {
	controller controller.ResourceController
}

func NewResourceApi(controller controller.ResourceController) *ResourceApi {
	return &ResourceApi{controller}
}


// Paths Information

// GetAws godoc
// @Summary List AWS Resources
// @Description Get all the AWS Resources
// @Tags AWS
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Resource
// @Failure 401 {object} dto.Response
// @Router /aws [get]
func (api *ResourceApi) GetAwsResources(ctx *gin.Context) {
	res, err := api.controller.GetAllAws()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{Response: err.Error()})
	} else {
		ctx.JSON(200, res)
	}
}

// GetAzure godoc
// @Summary List Azure Resources
// @Description Get all the Azure Resources
// @Tags Azure
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Resource
// @Failure 401 {object} dto.Response
// @Router /azure [get]
func (api *ResourceApi) GetAzureResources(ctx *gin.Context) {
	res, err := api.controller.GetAllAzure()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{Response: err.Error()})
	} else {
		ctx.JSON(200, res)
	}
}

// GetGcp godoc
// @Summary List GCP Resources
// @Description Get all the AWS Resources
// @Tags GCP
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Resource
// @Failure 401 {object} dto.Response
// @Router /gcp [get]
func (api *ResourceApi) GetGcpResources(ctx *gin.Context) {
	res, err := api.controller.GetAllGcp()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{Response: err.Error()})
	} else {
		ctx.JSON(200, res)
	}
}

// CreateAWS godoc
// @Summary Create new AWS Resource
// @Description Create a new AWS Resource
// @Tags AWS
// @Accept  json
// @Produce  json
// @Param resource body entity.Resource true "Create Resource"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /aws [post]
func (api *ResourceApi) AddAwsResource(ctx *gin.Context) {
	res, err := api.controller.AddAws(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{Response: err.Error()})
	} else {
		ctx.JSON(200, &dto.Response{Response: "Succesfully inserted AWS resource!", ID: res.InsertedID})
	}
}

// CreateAzure godoc
// @Summary Create new Azure Resource
// @Description Create a new Azure Resource
// @Tags Azure
// @Accept  json
// @Produce  json
// @Param resource body entity.Resource true "Create Resource"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /azure [post]
func (api *ResourceApi) AddAzureResource(ctx *gin.Context) {
	res, err := api.controller.AddAzure(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{Response: err.Error()})
	} else {
		ctx.JSON(200, &dto.Response{Response: "Succesfully inserted Azure resource!", ID: res.InsertedID})
	}
}

// CreateGcp godoc
// @Summary Create new GCP Resource
// @Description Create a new GCP Resource
// @Tags GCP
// @Accept  json
// @Produce  json
// @Param resource body entity.Resource true "Create Resource"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /gcp [post]
func (api *ResourceApi) AddGcpResource(ctx *gin.Context) {
	res, err := api.controller.AddGcp(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{Response: err.Error()})
	} else {
		ctx.JSON(200, &dto.Response{Response: "Succesfully inserted GCP resource!", ID: res.InsertedID})
	}
}