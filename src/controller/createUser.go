package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/julianyf/Desafio_BrandMonitor/src/configuration/rest_err"
	"github.com/julianyf/Desafio_BrandMonitor/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds, error=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(UserRequest)
}
