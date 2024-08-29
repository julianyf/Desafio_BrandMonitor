package converter

import (
	"github.com/julianyf/Desafio_BrandMonitor/src/model"
	"github.com/julianyf/Desafio_BrandMonitor/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
