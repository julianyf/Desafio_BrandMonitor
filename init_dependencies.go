package main

import (
	"github.com/julianyf/Desafio_BrandMonitor/src/controller"
	"github.com/julianyf/Desafio_BrandMonitor/src/model/repository"
	"github.com/julianyf/Desafio_BrandMonitor/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
