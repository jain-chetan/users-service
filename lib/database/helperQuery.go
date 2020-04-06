package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CheckUserExist checks whether User exist or not
func (dc *DBRepo) CheckUserExist(userID string) bool {

	collection := dc.DBClient.Database("shopping-cart").Collection("Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var count int64
	ID, err := primitive.ObjectIDFromHex(userID)

	count, err = collection.CountDocuments(ctx, bson.D{primitive.E{Key: "_id", Value: ID},
		primitive.E{Key: "IsDeleted", Value: false}})

	if count <= 0 || err != nil {
		if err != nil {
			log.Println("Error : ", err)
		}
		return false
	}

	return true
}
