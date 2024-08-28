package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/julianyf/Desafio_BrandMonitor/src/configuration/validation"
	"github.com/julianyf/Desafio_BrandMonitor/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(UserRequest)
}
