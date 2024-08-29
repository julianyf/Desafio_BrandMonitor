package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "github.com/julianyf/Desafio_BrandMonitor/src/configuration/database/mongodb"
	"github.com/julianyf/Desafio_BrandMonitor/src/configuration/database/mongodb"
	"github.com/julianyf/Desafio_BrandMonitor/src/controller"
	"github.com/julianyf/Desafio_BrandMonitor/src/controller/routes"
	"github.com/julianyf/Desafio_BrandMonitor/src/model/repository"
	"github.com/julianyf/Desafio_BrandMonitor/src/model/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying connect to database, error=%s \n",
			err.Error())
		return
	}

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
