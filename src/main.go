package main

import (
	"kinza/src/config"
	"kinza/src/docs"
	"kinza/src/router"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Gin Init
	app := gin.Default()

	// Config Init
	conf := config.LoadConfig("C:\\Progs\\kinza\\src\\config\\config.json")

	// Swagger Init
	docs.SwaggerInfo.BasePath = conf.BasePath

	// Router Init
	baseGroup := app.Group(conf.BasePath)
	{
		router.InitRoutes(baseGroup)
	}
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Servin Application
	app.Run(conf.Host + ":" + conf.Port)
}
