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
// @Accept  json
// @Produce  json
// @Param firstname formData string true "First Name"
// @Param lastname formData string true "Last Name"
// @Param graphtoken formData string true "Graph Token"
// @Success 200 {array} entity.AwsResource
// @Failure 401 {object} dto.Response
// @Router /aws [get]
func (api *ResourceApi) GetAwsResources(ctx *gin.Context) {
	ctx.JSON(200, api.controller.GetAllAws())
}

// GetAzure godoc
// @Summary List Azure Resources
// @Description Get all the Azure Resources
// @Accept  json
// @Produce  json
// @Param firstname formData string true "First Name"
// @Param lastname formData string true "Last Name"
// @Param graphtoken formData string true "Graph Token"
// @Success 200 {array} entity.AzureResource
// @Failure 401 {object} dto.Response
// @Router /azure [get]
func (api *ResourceApi) GetAzureResources(ctx *gin.Context) {
	ctx.JSON(200, api.controller.GetAllAzure())
}

// GetGcp godoc
// @Summary List GCP Resources
// @Description Get all the AWS Resources
// @Accept  json
// @Produce  json
// @Param firstname formData string true "First Name"
// @Param lastname formData string true "Last Name"
// @Param graphtoken formData string true "Graph Token"
// @Success 200 {array} entity.GcpResource
// @Failure 401 {object} dto.Response
// @Router /gcp [get]
func (api *ResourceApi) GetGcpResources(ctx *gin.Context) {
	ctx.JSON(200, api.controller.GetAllGcp())
}

// CreateAWS godoc
// @Summary Create new AWS Resource
// @Description Create a new AWS Resource
// @Accept  json
// @Produce  json
// @Param video body entity.AwsResource true "Create Resource"
// @Param firstname formData string true "First Name"
// @Param lastname formData string true "Last Name"
// @Param graphtoken formData string true "Graph Token"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /aws [post]
func (api *ResourceApi) AddAwsResource(ctx *gin.Context) {
	err := api.controller.AddAws(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}

// CreateAzure godoc
// @Summary Create new Azure Resource
// @Description Create a new Azure Resource
// @Accept  json
// @Produce  json
// @Param video body entity.AzureResource true "Create Resource"
// @Param firstname formData string true "First Name"
// @Param lastname formData string true "Last Name"
// @Param graphtoken formData string true "Graph Token"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /azure [post]
func (api *ResourceApi) AddAzureResource(ctx *gin.Context) {
	err := api.controller.AddAzure(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}

// CreateGcp godoc
// @Summary Create new GCP Resource
// @Description Create a new GCP Resource
// @Accept  json
// @Produce  json
// @Param video body entity.GcpResource true "Create Resource"
// @Param firstname formData string true "First Name"
// @Param lastname formData string true "Last Name"
// @Param graphtoken formData string true "Graph Token"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /gcp [post]
func (api *ResourceApi) AddGcpResource(ctx *gin.Context) {
	err := api.controller.AddGcp(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}