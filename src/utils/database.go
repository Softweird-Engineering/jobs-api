package utils

import (
	"context"
	"kinza/config"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once     sync.Once
	instance *mongo.Client
)

// Db
func Db() *mongo.Database {
	conf := config.Config()
	once.Do(func() {
		var err error
		instance, err = initClient(conf.Mongodb.Dsn, conf.Mongodb.Timeouts.Connection)
		if err != nil {
			log.WithFields(log.Fields{
				"Error": err,
			}).Error("can not init mongodb instance")
		}
	})
	return instance.Database(conf.Mongodb.Schema)
}

func initClient(dsn string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CloseClient(timeout int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	var err error
	if instance != nil {
		err = instance.Disconnect(ctx)
		if err != nil {
			log.WithFields(log.Fields{
				"Error": err,
			}).Error("can not disconnect instance")
		}
	} else {
		log.Error("instace is nil database instance closing before initialization")
	}
}
