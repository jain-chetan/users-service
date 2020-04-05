package database

import (
	"context"
	"log"
	"time"

	"github.com/jain-chetan/users-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetUserQuery - Query method to get the User details
func (dc *DBRepo) GetUserQuery(userID string) (model.User, error) {
	var user model.User

	collection := dc.DBClient.Database("shopping-cart").Collection("Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ID, errConversion := primitive.ObjectIDFromHex(userID)
	if errConversion != nil {
		return user, errConversion
	}

	err := collection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: ID},
		primitive.E{Key: "IsDeleted", Value: false}}).Decode(&user)

	if err != nil {
		log.Println("Error in getting user details ", err)
		return user, err
	}

	user.EncryptedPassword = ""

	return user, nil

}
