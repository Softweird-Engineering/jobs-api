package main

import (
	"errors"
	"fmt"
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
func InitApp() (*gin.Engine, error) {
	// Gin Init
	app := gin.New()
	app.Use(gin.Recovery())

	// Config Init
	conf := config.Config()
	log.Info("Config set up successfully!")

	// Logger init
	var logLevel log.Level
	if conf.Gin.Mode == "release" {
		app.Use(utils.Logger_JSON(conf.Log.Filename))
		logLevel = log.InfoLevel
	} else if conf.Gin.Mode == "debug" {
		gin.ForceConsoleColor()
		app.Use(gin.Logger())
		logLevel = log.DebugLevel
	} else if conf.Gin.Mode == "testing" {
		gin.ForceConsoleColor()
		logLevel = log.ErrorLevel
	} else {
		return nil, errors.New("gin.mode is invalid")
	}
	log.SetLevel(logLevel)

	// Router Init
	baseGroup := app.Group(conf.Server.BasePath)
	{
		router.InitRoutes(baseGroup)
	}

	// Swagger Init
	docs.SwaggerInfo.BasePath = conf.Server.BasePath
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return app, nil
}

func main() {
	conf := config.Config()

	app, err := InitApp()
	if err != nil {
		log.Fatal(err)
	}
	// Serving Application
	log.WithFields(log.Fields{
		"Conf": fmt.Sprintf("%#v\n", conf),
	}).Debug("Config for startup")
	log.Info("Starting application...")

	app.Run(conf.Server.Host + ":" + conf.Server.Port)
}
