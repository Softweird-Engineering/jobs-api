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

// InitApp ...
func InitApp() *gin.Engine {
	// Gin Init
	app := gin.New()
	app.Use(gin.Recovery())

	// Config Init
	conf := config.Config()
	log.Info("Config set up successfully!")

	// Logger init
	if conf.Gin.Mode == "release" {
		app.Use(utils.Logger_JSON(conf.Log.Filename))
	} else if conf.Gin.Mode == "debug" {
		gin.ForceConsoleColor()
		app.Use(gin.Logger())
	} else if conf.Gin.Mode == "testing" {
		gin.ForceConsoleColor()
	}

	// Router Init
	baseGroup := app.Group(conf.Server.BasePath)
	{
		router.InitRoutes(baseGroup)
	}

	// Swagger Init
	docs.SwaggerInfo.BasePath = conf.Server.BasePath
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return app
}

func main() {
	conf := config.Config()
	app := InitApp()
	// Serving Application
	log.Info("Starting application...")
	app.Run(conf.Server.Host + ":" + conf.Server.Port)
}
