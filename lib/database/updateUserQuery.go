package database

import (
	"context"
	"log"
	"time"

	"github.com/jain-chetan/users-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UpdateUserQuery - Database query method to update user
func (dc *DBRepo) UpdateUserQuery(userID string, user model.User) error {

	collection := dc.DBClient.Database("shopping-cart").Collection("Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//converting string ID into primitive hexadecimal object
	ID, errConversion := primitive.ObjectIDFromHex(userID)
	if errConversion != nil {
		return errConversion
	}

	_, errUpdate := collection.UpdateOne(ctx, bson.D{primitive.E{Key: "_id", Value: ID}},
		bson.D{primitive.E{Key: "$set", Value: user}})

	log.Println("Update filter and data ", bson.D{primitive.E{Key: "_id", Value: ID}},
		bson.D{primitive.E{Key: "$set", Value: user}})

	if errUpdate != nil {
		return errUpdate
	}

	return nil

}
