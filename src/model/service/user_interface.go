package service

import (
	"github.com/julianyf/Desafio_BrandMonitor/src/configuration/rest_err"
	"github.com/julianyf/Desafio_BrandMonitor/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
