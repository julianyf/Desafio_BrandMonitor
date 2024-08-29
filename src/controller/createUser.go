package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianyf/Desafio_BrandMonitor/src/configuration/logger"
	"github.com/julianyf/Desafio_BrandMonitor/src/configuration/validation"
	"github.com/julianyf/Desafio_BrandMonitor/src/controller/model/request"
	"github.com/julianyf/Desafio_BrandMonitor/src/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: UserRequest.Email,
		Name:  UserRequest.Name,
		Age:   UserRequest.Age,
	}

	logger.Info("User created sucessfully",
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, response)
}
