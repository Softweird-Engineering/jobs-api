package main

import (
	"context"
	"errors"
	"fmt"
	"kinza/config"
	"kinza/docs"
	"kinza/router"
	"kinza/utils"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	} else if conf.Gin.Mode == "test" {
		gin.ForceConsoleColor()
		logLevel = log.ErrorLevel
	} else {
		return nil, errors.New("gin.mode is invalid")
	}
	log.SetLevel(logLevel)

	// Db init
	db := utils.Db()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.Mongodb.Timeouts.Ping)*time.Second)
	err := db.Client().Ping(ctx, readpref.Primary())
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Error("troubles with ping")
	}
	cancel()

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

	// Inicialize
	app, err := InitApp()
	if err != nil {
		log.Fatal(err)
	}
	defer utils.CloseClient(conf.Mongodb.Timeouts.Connection)
	// Serving Application
	log.WithFields(log.Fields{
		"Conf": fmt.Sprintf("%#v\n", conf),
	}).Debug("Config for startup")
	log.Info("Starting application...")

	app.Run(conf.Server.Host + ":" + conf.Server.Port)
}
