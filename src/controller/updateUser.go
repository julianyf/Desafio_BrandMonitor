package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianyf/Desafio_BrandMonitor/src/configuration/logger"
	"github.com/julianyf/Desafio_BrandMonitor/src/configuration/rest_err"
	"github.com/julianyf/Desafio_BrandMonitor/src/configuration/validation"
	"github.com/julianyf/Desafio_BrandMonitor/src/controller/model/request"
	"github.com/julianyf/Desafio_BrandMonitor/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser controller",
		zap.String("journey", "updateUser"),
	)
	var UserRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"),
		)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(
		UserRequest.Name,
		UserRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("UpdateUser controller executed sucessfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	c.Status(http.StatusOK)

}
