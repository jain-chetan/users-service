package database

import (
	"context"
	"time"

	"github.com/jain-chetan/users-service/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateUserQuery - Database query method to create a user
func (dc *DBRepo) CreateUserQuery(user model.User) (model.CreateResponse, error) {
	var result model.CreateResponse
	emptyResponse := model.CreateResponse{}
	collection := dc.DBClient.Database("shopping-cart").Collection("Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return emptyResponse, err
	}

	result.ID = res.InsertedID.(primitive.ObjectID)

	return result, nil
}
