package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DeleteUserQuery - Deletes the particular user based on the ID passed
func (dc *DBRepo) DeleteUserQuery(userID string) error {
	collection := dc.DBClient.Database("shopping-cart").Collection("Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	setDeleteFlag := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "IsDeleted", Value: true},
		}},
	}

	_, errUpdate := collection.UpdateOne(ctx, bson.D{primitive.E{Key: "_id", Value: ID}}, setDeleteFlag)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
