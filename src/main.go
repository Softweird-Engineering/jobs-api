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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

var (
	mongoURI = "mongodb://localhost:27017"
)

func main() {
	conf := config.Config()

	// Db connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Buba:boba22@cluster0.sbhjz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	demoDB := client.Database("demo")
	catsCollection := demoDB.Collection("cats")
	cursor, err := catsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var cats []bson.M
	if err = cursor.All(ctx, &cats); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cats)

	// Inicialize
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
