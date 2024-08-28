package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "github.com/julianyf/Desafio_BrandMonitor/src/configuration/database/mongodb"
	"github.com/julianyf/Desafio_BrandMonitor/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// service := sevice.NewUserDomainService()

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
