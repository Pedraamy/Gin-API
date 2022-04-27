package middlewares

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedraamy/gin-api/dto"
)

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

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
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
			c.Next()
	}
}
