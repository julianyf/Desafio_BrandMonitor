package view

import (
	"github.com/julianyf/Desafio_BrandMonitor/src/controller/model/response"
	"github.com/julianyf/Desafio_BrandMonitor/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
