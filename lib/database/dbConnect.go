package database

import (
	"context"
	"log"
	"time"

	"github.com/jain-chetan/users-service/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DBRepo satisfies the interface by implementing all the methods
type DBRepo struct {
	DBClient *mongo.Client
}

//DBConnect Method to connect to DB
func (dc *DBRepo) DBConnect(config model.DBConfig) error {
	var err error
	//Set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//initialize DBClient
	dc.DBClient, err = mongo.Connect(ctx, options.Client().ApplyURI(config.User+"://"+config.Host+":"+config.Port))
	if err != nil {
		log.Printf("Unable to connect DB %v", err)
		return err
	}
	log.Printf("DB started at %s PORT", config.Port)
	return err
}
