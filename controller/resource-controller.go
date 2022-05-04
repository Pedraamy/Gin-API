package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pedraamy/gin-api/entity"
	"github.com/pedraamy/gin-api/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/go-playground/validator.v9"
)

type ResourceController interface {
	AddAws(c *gin.Context) (*mongo.InsertOneResult, error)
	AddAzure(c *gin.Context) error
	AddGcp(c *gin.Context) error
	/* DeleteAws(c *gin.Context) error
	DeleteAzure(c *gin.Context) error
	DeleteGcp(c *gin.Context) error */
	GetAllAws() ([]bson.D, error)
	GetAllAzure() ([]bson.D, error)
	GetAllGcp() ([]bson.D, error)
}

type controller struct {
	repo repository.ResourceRepo
	validate *validator.Validate
}

func NewController(repo repository.ResourceRepo) ResourceController {
	validate := validator.New()
	return &controller{repo: repo, validate: validate}
}

func (ctrl *controller) GetAllAws() ([]bson.D, error) {
	res, err := ctrl.repo.GetAllAws()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ctrl *controller) GetAllAzure() ([]bson.D, error) {
	res, err := ctrl.repo.GetAllAzure()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ctrl *controller) GetAllGcp() ([]bson.D, error) {
	res, err := ctrl.repo.GetAllGcp()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ctrl *controller) AddAws(c *gin.Context) (*mongo.InsertOneResult, error) {
	var resource entity.Resource
	err := c.ShouldBindJSON(&resource)
	if err != nil {
		return nil, err
	}
	err = ctrl.validate.Struct(resource)
	if err != nil {
		return nil, err
	}
	res, err := ctrl.repo.AddAws(resource)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ctrl *controller) AddAzure(c *gin.Context) error {
	var resource entity.Resource
	err := c.ShouldBindJSON(&resource)
	if err != nil {
		return err
	}
	err = ctrl.validate.Struct(resource)
	if err != nil {
		return err
	}
	err = ctrl.repo.AddAzure(resource)
	if err != nil {
		return err
	}
	return nil
}

func (ctrl *controller) AddGcp(c *gin.Context) error {
	var resource entity.Resource
	err := c.ShouldBindJSON(&resource)
	if err != nil {
		return err
	}
	err = ctrl.validate.Struct(resource)
	if err != nil {
		return err
	}
	err = ctrl.repo.AddGcp(resource)
	if err != nil {
		return err
	}
	return nil
}

/* func (ctrl *controller) DeleteAws(ctx *gin.Context) error {
	var video entity.Resource
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
} */