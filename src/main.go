package main

import (
	"kinza/config"
	"kinza/docs"
	"kinza/router"
	"kinza/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Gin Init
	app := gin.New()
	app.Use(gin.Recovery())

	// Config Init
	conf := config.Config(".")
	log.Info("Config set up successfully!")

	// Logger init
	app.Use(utils.Logger_JSON(conf.Log.Filename, conf.Gin.Mode))

	// Router Init
	baseGroup := app.Group(conf.Server.BasePath)
	{
		router.InitRoutes(baseGroup)
	}

	// Swagger Init
	docs.SwaggerInfo.BasePath = conf.Server.BasePath
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Serving Application
	log.Info("Starting application...")
	app.Run(conf.Server.Host + ":" + conf.Server.Port)
}
