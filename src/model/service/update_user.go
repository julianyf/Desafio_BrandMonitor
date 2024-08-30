package service

import (
	"github.com/julianyf/Desafio_BrandMonitor/src/configuration/logger"
	"github.com/julianyf/Desafio_BrandMonitor/src/configuration/rest_err"
	"github.com/julianyf/Desafio_BrandMonitor/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUser model",
		zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateUser"))
		return err
	}

	logger.Info(
		"CreatedUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))
	return nil
}
