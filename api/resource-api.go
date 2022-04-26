package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/dto"
)

type ResourceApi struct {
	loginController controller.LoginController
	videoController controller.VideoController
}

func NewVideoAPI(loginController controller.LoginController,
	videoController controller.VideoController) *ResourceApi {
	return &ResourceApi{
		loginController: loginController,
		videoController: videoController,
	}
}

func GetResponseBody(url string, headerKey string, headerValue string) (map[string]interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(headerKey, headerValue)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Paths Information

// Authenticate godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param firstname formData string true "First Name"
// @Param lastname formData string true "Last Name"
// @Param graphtoken formData string true "Graph Token"
// @Success 200 {object} dto.JWT
// @Failure 401 {object} dto.Response
// @Router /auth/token [post]
func (api *ResourceApi) Authenticate(c *gin.Context) {
	headerVal := c.GetHeader("Authorization")
	if len(headerVal) == 0 {
		resp := "Authorization header is incorrect or not provided. You must use 'Authorization' as header key."
		c.AbortWithStatusJSON(401, resp)
		return
	}
	headerVal = "Bearer " + headerVal
	url := "https://graph.microsoft.com/v1.0/me/"
	res, err := GetResponseBody(url, "Authorization", headerVal)
	if err != nil {
		resp := "Error occured while attempting to access credentials."
		c.AbortWithStatusJSON(400, resp)
		return
	}
	accountName := fmt.Sprintf("%s", res["userPrincipalName"])
	var company string = ""
	for i := 0; i < len(accountName); i++ {
		if accountName[i] == 64 {
			company = accountName[i+1:]
			break
		}
	}
	if company != "zuoracloudeng.onmicrosoft.com" {
		resp := "Invalid credentials."
		c.AbortWithStatusJSON(401, resp)
		return
	}
	name := fmt.Sprintf("%s", res["displayName"])

	var credentials dto.Name
	err = c.ShouldBind(&credentials)
	if err != nil {
		resp := "Error occured while attempting to access credentials."
		c.AbortWithStatusJSON(400, resp)
		return
	}

	comp := credentials.Firstname + credentials.Lastname
	if comp != name {
		resp := "Name does not match credentials."
		c.AbortWithStatusJSON(401, resp)
		return
	}

	token := api.loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, &dto.JWT{
			Token: token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, &dto.Response{
			Message: "Not Authorized",
		})
	}
}

// GetVideos godoc
// @Security bearerAuth
// @Summary List existing videos
// @Description Get all the existing videos
// @Tags videos,list
// @Accept  json
// @Produce  json
// @Param firstname formData string true "First Name"
// @Param lastname formData string true "Last Name"
// @Param graphtoken formData string true "Graph Token"
// @Success 200 {array} entity.AwsResource
// @Failure 401 {object} dto.Response
// @Router /videos [get]
func (api *ResourceApi) GetAwsResources(ctx *gin.Context) {
	ctx.JSON(200, api.videoController.FindAllAws())
}

// GetVideos godoc
// @Security bearerAuth
// @Summary List existing videos
// @Description Get all the existing videos
// @Tags videos,list
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Video
// @Failure 401 {object} dto.Response
// @Router /videos [get]
func (api *ResourceApi) GetAzureResources(ctx *gin.Context) {
	ctx.JSON(200, api.videoController.FindAllAzure())
}

// GetVideos godoc
// @Security bearerAuth
// @Summary List existing videos
// @Description Get all the existing videos
// @Tags videos,list
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Video
// @Failure 401 {object} dto.Response
// @Router /videos [get]
func (api *ResourceApi) GetGcpResources(ctx *gin.Context) {
	ctx.JSON(200, api.videoController.FindAllGcp())
}

// CreateVideo godoc
// @Security bearerAuth
// @Summary Create new videos
// @Description Create a new video
// @Tags videos,create
// @Accept  json
// @Produce  json
// @Param video body entity.Video true "Create video"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos [post]
func (api *ResourceApi) AddAwsResource(ctx *gin.Context) {
	err := api.videoController.Save(ctx)
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

// CreateVideo godoc
// @Security bearerAuth
// @Summary Create new videos
// @Description Create a new video
// @Tags videos,create
// @Accept  json
// @Produce  json
// @Param video body entity.Video true "Create video"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos [post]
func (api *ResourceApi) AddAzureResource(ctx *gin.Context) {
	err := api.videoController.Save(ctx)
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

// CreateVideo godoc
// @Security bearerAuth
// @Summary Create new videos
// @Description Create a new video
// @Tags videos,create
// @Accept  json
// @Produce  json
// @Param video body entity.Video true "Create video"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos [post]
func (api *ResourceApi) AddGcpResource(ctx *gin.Context) {
	err := api.videoController.Save(ctx)
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

// UpdateVideo godoc
// @Security bearerAuth
// @Summary Update videos
// @Description Update a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Param video body entity.Video true "Update video"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos/{id} [put]
func (api *VideoApi) UpdateVideo(ctx *gin.Context) {
	err := api.videoController.Update(ctx)
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

// DeleteVideo godoc
// @Security bearerAuth
// @Summary Remove videos
// @Description Delete a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos/{id} [delete]
func (api *VideoApi) DeleteVideo(ctx *gin.Context) {
	err := api.videoController.Delete(ctx)
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
