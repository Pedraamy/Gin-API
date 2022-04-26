package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/repository"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
	"gopkg.in/go-playground/validator.v9"
)

type ResourceController interface {
	AddAws(c *gin.Context) error
	AddAzure(c *gin.Context) error
	AddGcp(c *gin.Context) error
	DeleteAws(c *gin.Context) error
	DeleteAzure(c *gin.Context) error
	DeleteGcp(c *gin.Context) error
	GetAllAws() []entity.AwsResource
	GetAllAzure() []entity.AzureResource
	GetAllGcp() []entity.GcpResource
}

type controller struct {
	repo repository.ResourceRepo
}

var validate *validator.Validate

func NewController(repo repository.ResourceRepo) *controller {
	return &controller{repo: repo}
}

func (ctrl *controller) GetAllAws() []entity.AwsResource {
	return ctrl.repo.GetAllAws()
}

func (ctrl *controller) GetAllAzure() []entity.AzureResource {
	return ctrl.repo.GetAllAzure()
}

func (ctrl *controller) GetAllGcp() []entity.GcpResource {
	return ctrl.repo.GetAllGcp()
}

func (ctrl *controller) AddAws(c *gin.Context) error {
	var resource entity.AwsResource
	err := c.ShouldBindJSON(&resource)
	if err != nil {
		return err
	}
	err = validate.Struct(resource)
	if err != nil {
		return err
	}
	ctrl.repo.AddAws(resource)
	return nil
}

func (ctrl *controller) AddAzure(c *gin.Context) error {
	var resource entity.AzureResource
	err := c.ShouldBindJSON(&resource)
	if err != nil {
		return err
	}
	err = validate.Struct(resource)
	if err != nil {
		return err
	}
	ctrl.repo.AddAzure(resource)
	return nil
}

func (ctrl *controller) AddGcp(c *gin.Context) error {
	var resource entity.GcpResource
	err := c.ShouldBindJSON(&resource)
	if err != nil {
		return err
	}
	err = validate.Struct(resource)
	if err != nil {
		return err
	}
	ctrl.repo.AddGcp(resource)
	return nil
}

func (ctrl *controller) DeleteAws(ctx *gin.Context) error {
	var video entity.AwsResource
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	ctrl.repo.DeleteAws(video)
	return nil
}

func (ctrl *controller) DeleteAzure(ctx *gin.Context) error {
	var video entity.AzureResource
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	ctrl.repo.DeleteAzure(video)
	return nil
}

func (ctrl *controller) DeleteGcp(ctx *gin.Context) error {
	var video entity.GcpResource
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	ctrl.repo.DeleteGcp(video)
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id

	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Update(video)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var video entity.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	c.service.Delete(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
